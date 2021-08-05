package cmd

import (
	"os"
	"fmt"
	"github.com/spf13/cobra"
	"local.packages/expedition3gpp"
)

var rootCmd = &cobra.Command{}

func Execute() {
	err := rootCmd.Execute()
    if err != nil {
        os.Exit(0)
    }
}

func init() {
	rootCmd.Use = "expedition3gpp"
	rootCmd.Short = "Download the 3GPP document"

	var version bool
	rootCmd.Flags().BoolVarP(&version, "version", "v", false, "display version")

	rootCmd.RunE = func(cmd *cobra.Command, args []string) error {
		// Assumption to be executed only the first time.
		// Runs if the config file does not exist.
		if expedition3gpp.ExistInitConfig() {
			initConfig := expedition3gpp.InitConfig{
				StrageLocation:     "HOMEDIR",
				CacheEnable:        true,
				CacheRetentionTime: 14400,
				CacheLocation:      "HOMEDIR",
			}
		
			expedition3gpp.InitializeConfig(&initConfig)
			fmt.Println("Create config file")
		}

		if version {
			fmt.Println("version: 1.0.0")
			os.Exit(0)
		}
		return rootCmd.Help()
	}
}
