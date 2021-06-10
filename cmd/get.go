package cmd

import (
	log "github.com/justsimplify/totp-cli/logger"
	"github.com/justsimplify/totp-cli/modules"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var getCmd = &cobra.Command{
	Use:   modules.GetUse,
	Short: "Get secret (totp)",
	PreRun: func(cmd *cobra.Command, args []string) {
		_ = viper.BindPFlags(cmd.Flags())
	},
	Run: getFunc,
}

func init() {
	modules.AddFlags(getCmd)
}

func init() {
	rootCmd.AddCommand(getCmd)
}

func getFunc(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		log.Fatal("error occurred")
		return
	}

	m, ok := modules.Modules[args[0]]
	if !ok {
		log.Fatalf("functionality for '%v' doesn't exist", args[0])
		return
	}

	cModule, ok := m().(modules.GetModule)

	if !ok {
		log.Fatalf("can't perform get action on '%v'", args[0])
		return
	}

	_, err := cModule.Get(cmd, args[1:])

	if err != nil {
		log.Fatal(err.Error())
		return
	}
}
