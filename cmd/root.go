package cmd

import (
	"os"
	"fmt"
	"github.com/spf13/cobra"
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
		if version {
			fmt.Println("version: 1.0.0")
			os.Exit(0)
		}
		return rootCmd.Help()
	}
}
