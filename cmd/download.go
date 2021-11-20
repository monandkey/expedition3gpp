package cmd

import (
	"errors"
	"os"

	"github.com/spf13/cobra"
	"local.packages/expedition3gpp"
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
			return errors.New("Specify the document number.\n")
		}

		config := expedition3gpp.Config{
			DocumentNumber:  params.documentNumber,
			DocumentVersion: params.documentVersion,
			OutputPath:      params.outputPath,
			Cache:           params.cache,
		}

		err := expedition3gpp.RunExpedition3gpp(&config)
		if err != nil {
			return err
		}
		return nil
	}
	rootCmd.AddCommand(downloadCmd)
}
