package totp

import (
	"time"

	"github.com/pquerna/otp"
	"github.com/pquerna/otp/totp"
)

func (t *TOTP) GenerateCurrentTOTP(secret, issuer string, digits int) (string, error) {
	return totp.GenerateCodeCustom(secret, time.Now(), totp.ValidateOpts{
		Digits: otp.Digits(digits),
	})
}
