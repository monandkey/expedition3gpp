package cmd

import (
	"os"
	"strconv"
	"strings"

	"github.com/monandkey/expedition3gpp/pkg/config"
	"github.com/spf13/cobra"
)

func init() {
	initCmd := &cobra.Command{}
	initCmd.Use = "init"
	initCmd.Short = "Create the config file"

	initConfig := configParams{
		strageLocation:     "HOMEDIR",
		cacheEnable:        true,
		cacheRetentionTime: 14400,
		cacheLocation:      "HOMEDIR",
	}

	initCmd.Flags().StringVarP(&initConfig.strageLocation, "strage-location", "s", initConfig.strageLocation, "Specify the location to save the config.\nwindows -> C:\\Users\\testuser\nlinux   -> /home/testuser\n")
	initCmd.Flags().BoolVarP(&initConfig.cacheEnable, "cache-enable", "e", initConfig.cacheEnable, "Enable or disable the cache\ntrue  -> enable\nfalse -> disable\n")
	initCmd.Flags().IntVarP(&initConfig.cacheRetentionTime, "cache-retention-time", "r", initConfig.cacheRetentionTime, "Specify the validity period for saving the cache.\n[0...4294967295]\n")
	initCmd.Flags().StringVarP(&initConfig.cacheLocation, "cache-location", "l", initConfig.cacheLocation, "Specify the location to save the cache.\nwindows -> C:\\Users\\testuser\nlinux   -> /home/testuser\n")

	initCmd.RunE = func(cmd *cobra.Command, args []string) error {
		// If false is specified, it overrides the boolean value.
		for i, v := range os.Args {
			if strings.Contains(v, "cache-enable") && os.Args[i+1] == "false" {
				initConfig.cacheEnable, _ = strconv.ParseBool(os.Args[i+1])
			}
		}
		c := config.SelectConfigUser()
		c.SetParams(
			initConfig.strageLocation,
			initConfig.cacheEnable,
			initConfig.cacheRetentionTime,
			initConfig.strageLocation,
		)
		if err := c.Write(); err != nil {
			return err
		}

		return nil
	}
	rootCmd.AddCommand(initCmd)
}
