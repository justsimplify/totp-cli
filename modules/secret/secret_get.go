package secret

import (
	"bytes"
	"fmt"
	"image"
	"net/url"
	"os"
	"strings"

	_ "image/jpeg"
	_ "image/png"

	log "github.com/justsimplify/totp-cli/logger"
	"github.com/justsimplify/totp-cli/models/errormodel"
	"github.com/justsimplify/totp-cli/models/secretmodel"
	"github.com/justsimplify/totp-cli/modules/totp"
	"github.com/makiuchi-d/gozxing"
	"github.com/makiuchi-d/gozxing/qrcode"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func (secret *Secret) Get(cmd *cobra.Command, args []string) (interface{}, error) {
	filePath := viper.GetString("filepath")
	uri := viper.GetString("uri")
	if uri != "" && filePath != "" {
		return nil, fmt.Errorf("only one of URI or File Path should be provided")
	}

	digits := viper.GetInt("digits")

	if filePath != "" {
		err := secret.ValidateQRFilePathExtension(filePath)
		if err != nil {
			return nil, err
		}

		result, err := secret.ReadQR(filePath)
		if err != nil {
			return nil, err
		}

		response, err := secret.ValidateURI(result.String())
		if err != nil {
			return nil, err
		}

		tp := totp.TOTP{}
		code, err := tp.GenerateCurrentTOTP(response.Secret, response.Issuer, digits)
		if err != nil {
			return nil, err
		}

		log.Infof("TOTP now is %v\n", code)

		return result, nil
	}

	if uri != "" {
		response, err := secret.ValidateURI(uri)
		if err != nil {
			return nil, err
		}

		tp := totp.TOTP{}
		code, err := tp.GenerateCurrentTOTP(response.Secret, response.Issuer, digits)
		if err != nil {
			return nil, err
		}

		log.Infof("TOTP now is %v\n", code)
	}

	return nil, nil
}

func (secret *Secret) ValidateQRFilePathExtension(filePath string) error {
	if !strings.HasSuffix(filePath, ".png") && !strings.HasSuffix(filePath, ".jpg") && !strings.HasSuffix(filePath, ".jpeg") {
		return fmt.Errorf("image should be of PNG, JPG or JPEG format")
	}
	return nil
}

func (secret *Secret) ReadQR(filePath string) (*gozxing.Result, error) {
	b, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	img, _, err := image.Decode(bytes.NewReader(b))
	if err != nil {
		return nil, err
	}

	bmp, _ := gozxing.NewBinaryBitmapFromImage(img)

	// decode image
	qrReader := qrcode.NewQRCodeReader()
	result, err := qrReader.Decode(bmp, nil)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (qr *Secret) ValidateURI(uri string) (*secretmodel.SecretResponse, error) {
	uri, err := url.PathUnescape(uri)
	if err != nil {
		return nil, err
	}

	uriParsed, err := url.Parse(uri)
	if err != nil {
		return nil, err
	}

	er := &errormodel.MultiError{}
	if uriParsed.Scheme != "otpauth" {
		er.Append(fmt.Errorf("scheme should be 'otpauth'"))
	}

	if uriParsed.Host != "totp" {
		er.Append(fmt.Errorf("hostname should be 'totp'"))
	}

	if uriParsed.Query().Get("secret") == "" {
		er.Append(fmt.Errorf("no secret found in url"))
	}

	if uriParsed.Query().Get("issuer") == "" {
		er.Append(fmt.Errorf("no issuer found in url"))
	}

	if !er.CheckIfNoError() {
		return nil, er
	}

	res := secretmodel.SecretResponse{
		Issuer: uriParsed.Query().Get("issuer"),
		Secret: uriParsed.Query().Get("secret"),
	}

	return &res, nil
}
