package cmd

import (
	"errors"
	"os"

	"github.com/monandkey/expedition3gpp/internal/pkg/expedition"
	"github.com/spf13/cobra"
)

func init() {
	downloadCmd := &cobra.Command{}
	downloadCmd.Use = "download"
	downloadCmd.Short = "Download the 3GPP documentation."

	params := params{
		documentNumber:  "",
		documentVersion: "",
		outputPath:      "",
		cache:           false,
	}

	downloadCmd.Flags().StringVarP(&params.documentNumber, "document-number", "n", params.documentNumber, "3GPP Document Number")
	downloadCmd.Flags().StringVarP(&params.documentVersion, "document-version", "v", params.documentVersion, "3GPP Document Version")
	downloadCmd.Flags().StringVarP(&params.outputPath, "output-path", "o", params.outputPath, "Specify the output location of the file")
	downloadCmd.Flags().BoolVarP(&params.cache, "no-cache", "c", params.cache, "Not using cache")

	downloadCmd.RunE = func(cmd *cobra.Command, args []string) error {
		if len(os.Args) < 3 {
			return downloadCmd.Help()
		}

		if params.documentNumber == "" {
			return errors.New("specify the document number")
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

		if err := actor.Download(); err != nil {
			return err
		}

		if err := actor.Cache(); err != nil {
			return err
		}
		return nil
	}
	rootCmd.AddCommand(downloadCmd)
}
