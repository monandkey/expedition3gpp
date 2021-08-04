package cmd

import (
	"os"
	"fmt"
	"github.com/spf13/cobra"
	"local.packages/expedition3gpp"
)

var initCmd = &cobra.Command{
		Use:   "init [OPTION]...",
		Short: "Create the config file",
		Run:    func(cmd *cobra.Command, args []string)  {
			if expedition3gpp.ExistInitConfig() {
				expedition3gpp.InitializeConfig()
				fmt.Println("Create confgi file")
				os.Exit(0)
			}
		},
}

func init() {	
	rootCmd.AddCommand(initCmd)

	initConfig := expedition3gpp.GetConfigParameter()

	initCmd.Flags().StringVar(&initConfig.StrageLocation, "strage-location", initConfig.StrageLocation, "")
	initCmd.Flags().BoolVar(&initConfig.CacheEnable, "cache-enable", initConfig.CacheEnable, "")
	initCmd.Flags().IntVar(&initConfig.CacheRetentionTime, "cache-retention-time", initConfig.CacheRetentionTime, "")
	initCmd.Flags().StringVar(&initConfig.CacheLocation, "cache-location", initConfig.CacheLocation, "")

}
