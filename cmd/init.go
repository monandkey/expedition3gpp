package cmd

import (
	"os"
	"fmt"
	"strings"
	"strconv"
	"github.com/spf13/cobra"
	"local.packages/expedition3gpp"
)

func init() {
	initCmd := &cobra.Command{}
	initCmd.Use = "init"
	initCmd.Short = "Create the config file"

	initConfig := expedition3gpp.InitConfig{
		StrageLocation:     "HOMEDIR",
		CacheEnable:        true,
		CacheRetentionTime: 14400,
		CacheLocation:      "HOMEDIR",
	}

	initCmd.Flags().StringVar(&initConfig.StrageLocation, "strage-location", initConfig.StrageLocation, "Specify the location to save the config.\nwindows -> C:\\Users\\testuser\nlinux   -> /home/testuser\n")
	initCmd.Flags().BoolVar(&initConfig.CacheEnable, "cache-enable", initConfig.CacheEnable, "Enable or disable the cache\ntrue  -> enable\nfalse -> disable\n")
	initCmd.Flags().IntVar(&initConfig.CacheRetentionTime, "cache-retention-time", initConfig.CacheRetentionTime, "Specify the validity period for saving the cache.\n[0...4294967295]\n")
	initCmd.Flags().StringVar(&initConfig.CacheLocation, "cache-location", initConfig.CacheLocation, "Specify the location to save the cache.\nwindows -> C:\\Users\\testuser\nlinux   -> /home/testuser\n")

	initCmd.RunE = func(cmd *cobra.Command, args []string) error {
		// If false is specified, it overrides the boolean value.
		for i, v := range os.Args {
			if strings.Contains(v, "cache-enable") && os.Args[i+1] == "false" {
				initConfig.CacheEnable, _ = strconv.ParseBool(os.Args[i+1])
			}
		}

		// If there is no config file
		if expedition3gpp.ExistInitConfig() {
			expedition3gpp.InitializeConfig(&initConfig)
			fmt.Println("Create config file")
			os.Exit(0)
		
		} else {
			var a string
			fmt.Printf("overwrite ? y or n: ")
			fmt.Scan(&a)

			// If you want to overwrite the config file
			if a == "y" {
				expedition3gpp.InitializeConfig(&initConfig)
				fmt.Println("Overwrite!!")
			}
		}
		return nil
	}
	rootCmd.AddCommand(initCmd)
}
