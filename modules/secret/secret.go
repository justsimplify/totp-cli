package secret

import "github.com/justsimplify/totp-cli/modules"

// Secret for qr
type Secret struct {
}

func init() {
	modules.ModuleAdd("secret", func() interface{} {
		return &Secret{}
	})
}
