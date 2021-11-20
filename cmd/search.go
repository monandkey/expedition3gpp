package cmd

import (
	"errors"
	"os"

	"github.com/spf13/cobra"
	"local.packages/expedition3gpp"
)

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

	searchCmd.Flags().StringVarP(&params.documentNumber, "document-number", "n", params.documentNumber, "3GPP Document Number")
	searchCmd.Flags().StringVarP(&params.documentVersion, "document-version", "v", params.documentVersion, "3GPP Document Version")
	searchCmd.Flags().BoolVarP(&params.cache, "no-cache", "c", params.cache, "Not using cache")

	searchCmd.RunE = func(cmd *cobra.Command, args []string) error {
		if len(os.Args) < 3 {
			return searchCmd.Help()
		}

		if params.documentNumber == "" {
			return errors.New("Specify the document.\n")
		}

		config := expedition3gpp.Config{
			DocumentNumber:  params.documentNumber,
			DocumentVersion: params.documentVersion,
			OutputPath:      params.outputPath,
			Cache:           params.cache,
		}

		err := expedition3gpp.SearchExpedition3gpp(&config)
		if err != nil {
			return err
		}
		return nil
	}
	rootCmd.AddCommand(searchCmd)
}
