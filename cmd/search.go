package cmd

import (
	"errors"
	"os"

	"github.com/monandkey/expedition3gpp/internal/pkg/expedition"
	"github.com/spf13/cobra"
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
			return errors.New("specify the document")
		}

		actor := expedition.SelectUser()
		actor.SetParams(
			params.documentNumber,
			params.documentVersion,
			params.outputPath,
			params.cache,
		)
		if err := actor.Search(); err != nil {
			return err
		}

		if err := actor.Cache(); err != nil {
			return err
		}
		return nil
	}
	rootCmd.AddCommand(searchCmd)
}
