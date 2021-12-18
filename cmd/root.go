package cmd

import (
	"os"

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
	rootCmd.Version = "1.0.1"
	rootCmd.SilenceUsage = true
	rootCmd.RunE = func(cmd *cobra.Command, args []string) error {
		return rootCmd.Help()
	}
}
