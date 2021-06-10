package secret

import (
	"github.com/justsimplify/totp-cli/modules"
	"github.com/spf13/cobra"
)

// Flags - Add flags
func (qr *Secret) Flags(cmd *cobra.Command) {
	rootCommand := cmd.Use
	switch rootCommand {
	case modules.GetUse:
		addFilePath(cmd)
		addURI(cmd)
		addDigits(cmd)
	}
}

func addFilePath(cmd *cobra.Command) {
	cmd.PersistentFlags().StringP("filepath", "f", "", "path to file")
}

func addURI(cmd *cobra.Command) {
	cmd.PersistentFlags().String("uri", "", "TOTP URI")
}

func addDigits(cmd *cobra.Command) {
	cmd.PersistentFlags().IntP("digits", "d", 6, "digits for the otp")
}
