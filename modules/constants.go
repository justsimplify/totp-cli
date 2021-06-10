package modules

import "github.com/mitchellh/go-homedir"

// HomePath for config
var HomePath, _ = homedir.Dir()

// Constants for cmd
const (
	GetUse      = "get"
	RootCmd     = "totp-cli"
	CfgFilePath = ".totp-cli"
	CfgFileName = "config.yml"
)
