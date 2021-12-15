package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}

func init() {
	rootCmd.Use = "expedition3gpp"
	rootCmd.Short = "Download the 3GPP document"
	rootCmd.Version = "1.0.0"
	rootCmd.SilenceUsage = true
	rootCmd.RunE = func(cmd *cobra.Command, args []string) error {
		return rootCmd.Help()
	}
}
