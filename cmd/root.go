package cmd

import (
	"os"
	"path"

	log "github.com/justsimplify/totp-cli/logger"
	"github.com/justsimplify/totp-cli/modules"
	_ "github.com/justsimplify/totp-cli/modules/secret"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:   modules.RootCmd,
	Short: "CLI for lazy",
}

func init() {
	cobra.OnInitialize(initConfig)
}

func initConfig() {
	cfgFilePath := path.Join(modules.HomePath, modules.CfgFilePath)

	err := os.MkdirAll(cfgFilePath, 0777)
	if err != nil {
		log.Fatal(err.Error())
	}

	f, err := os.OpenFile(cfgFilePath+"/"+modules.CfgFileName, os.O_RDONLY|os.O_CREATE, 0777)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer f.Close()

	viper.AddConfigPath(cfgFilePath)
	viper.SetConfigName("config")
	viper.SetDefault("colorize", true)

	viper.AutomaticEnv()
	err = viper.ReadInConfig()
	if err != nil {
		log.Fatal(err.Error())
	}
}

func Run() error {
	err := rootCmd.Execute()
	if err != nil {
		return err
	}
	return nil
}
