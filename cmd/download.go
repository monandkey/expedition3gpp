package cmd

import (
	"os"
	"fmt"
	"errors"
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

	downloadCmd.Flags().StringVar(&params.documentNumber, "document-number", params.documentNumber, "3GPP Document Number")
	downloadCmd.Flags().StringVar(&params.documentVersion, "document-version", params.documentVersion, "3GPP Document Version")
	downloadCmd.Flags().StringVar(&params.outputPath, "output-path", params.outputPath, "Specify the output location of the file")
	downloadCmd.Flags().BoolVar(&params.cache, "no-cache", params.cache, "Not using cache")

	downloadCmd.RunE = func(cmd *cobra.Command, args []string) error {
		if len(os.Args) < 3 {
			return errors.New("The argument is missing.")
		}

		if params.documentNumber == "" {
			return errors.New("Specify the document number.")
		}

		config := expedition3gpp.Config{
			DocumentNumber:  params.documentNumber,
			DocumentVersion: params.documentVersion,
			OutputPath:      params.outputPath,
			Cache:           params.cache,
		}

		err := expedition3gpp.RunExpedition3gpp(&config)
		if err != nil {
			fmt.Println(os.Stderr, err)
			os.Exit(0)
		}
		return nil
	}
	rootCmd.AddCommand(downloadCmd)
}