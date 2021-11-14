package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"local.packages/expedition3gpp"
)

type params struct {
	documentNumber  string
	documentVersion string
	outputPath      string
	cache           bool
}

func init() {
	searchCmd := &cobra.Command{}
	searchCmd.Use = "search"
	searchCmd.Short = "Search for 3GPP documentation."

	params := params{
		documentNumber:  "",
		documentVersion: "",
		outputPath:      "",
		cache:           false,
	}

	searchCmd.Flags().StringVar(&params.documentNumber, "document-number", params.documentNumber, "3GPP Document Number")
	searchCmd.Flags().StringVar(&params.documentVersion, "document-version", params.documentVersion, "3GPP Document Version")
	searchCmd.Flags().BoolVar(&params.cache, "no-cache", params.cache, "Not using cache")

	searchCmd.RunE = func(cmd *cobra.Command, args []string) error {
		if len(os.Args) < 3 {
			return errors.New("The argument is missing.")
		}

		if params.documentNumber == "" {
			return errors.New("Specify the document.")
		}

		config := expedition3gpp.Config{
			DocumentNumber:  params.documentNumber,
			DocumentVersion: params.documentVersion,
			OutputPath:      params.outputPath,
			Cache:           params.cache,
		}

		err := expedition3gpp.SearchExpedition3gpp(&config)
		if err != nil {
			fmt.Println(os.Stderr, err)
			os.Exit(0)
		}
		return nil
	}
	rootCmd.AddCommand(searchCmd)
}
