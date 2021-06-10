package modules

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func PurgeConfigs() {
	viper.Set("filepath", "")
	viper.Set("uri", "")
	viper.Set("digits", 6)
	_ = viper.WriteConfig()
}

func AddFlags(cmd *cobra.Command) {
	if len(os.Args) > 2 {
		m := os.Args[2]

		module, err := Get(m)
		if err != nil {
			return
		}
		flagModule, ok := module.(FlagModule)
		if !ok {
			return
		}
		flagModule.Flags(cmd)
	}
}

func GetFilePath() (string, error) {
	path := viper.GetString("filepath")
	if path == "" {
		return "", nil
	}

	if path == "-" {
		return path, nil
	}

	return path, nil
}
