package totp

import "github.com/justsimplify/totp-cli/modules"

// TOTP for totp
type TOTP struct {
}

func init() {
	modules.ModuleAdd("qr", func() interface{} {
		return &TOTP{}
	})
}
