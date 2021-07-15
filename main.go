package main

import (
	"os"
	"fmt"

	"github.com/spf13/cobra"
	"local.packages/expedition3gpp"
)

type params struct {
	url string
	documentNumber string
	documentVersion string
	outputPath string
}

// --------------------------------------------------
// Main
// --------------------------------------------------
func main() {
	params := params{
		url:            "default",
		documentNumber: "default",
		documentVersion: "default",	
		outputPath:      "./",
	}

	cmd := &cobra.Command{}
	cmd.Use = "expedition3gpp [OPTIONS] DOCUMENT_NUMBER DOCUMENT_VERSION OUTPUT_PATH"
	cmd.Short = "Download the 3GPP document"

	cmd.Flags().StringVarP(&params.url, "url", "u", params.url, "3GPP Doument URL")
	cmd.Flags().StringVarP(&params.documentNumber, "document-number", "n", params.documentNumber, "3GPP Document Number")
	cmd.Flags().StringVarP(&params.documentVersion, "document-version", "v", params.documentVersion, "3GPP Document Version")
	cmd.Flags().StringVarP(&params.outputPath, "output-path", "o", params.outputPath, "Where you want the output to go")

	cmd.RunE = func(cmd *cobra.Command, args []string) error {
		if params.url == "default" && params.documentNumber == "default" {
			return cmd.Help()
		}

		config := expedition3gpp.Config{
			Url:             params.url,
			DocumentNumber:  params.documentNumber,
			DocumentVersion: params.documentVersion,
			OutputPath:      params.outputPath,
		}

		err := expedition3gpp.RunExpedition3gpp(&config)
		if err != nil {
			fmt.Println(os.Stderr, err)
			os.Exit(1)
		}

		return nil
	}

	if err := cmd.Execute(); err != nil {
		fmt.Println(os.Stderr, err)
		os.Exit(1)
	}
}