package cmd

import (
	"os"
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"local.packages/expedition3gpp"
)

type params struct {
	documentNumber  string
	documentVersion string
	outputPath      string
	cache           bool
}

var rootCmd = &cobra.Command{
	Use:   "expedition3gpp",
	Short: "Download the 3GPP document",
}

func Execute() {
	err := rootCmd.Execute()
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
}

func init() {
	params := params{
		documentNumber:  "",
		documentVersion: "",	
		outputPath:      "current",
		cache:           false,
	}

	rootCmd.Flags().StringVar(&params.documentNumber, "document-number", params.documentNumber, "3GPP Document Number")
	rootCmd.Flags().StringVar(&params.documentVersion, "document-version", params.documentVersion, "3GPP Document Version")
	rootCmd.Flags().StringVar(&params.outputPath, "output-path", params.outputPath, "Where you want the output to go")
	rootCmd.Flags().BoolVar(&params.cache, "no-cache", params.cache, "Not using cache")


	rootCmd.RunE = func(cmd *cobra.Command, args []string) error {
		for _, v := range os.Args {
			if strings.Contains(v, "init") || expedition3gpp.ExistInitConfig() {
				expedition3gpp.InitializeConfig()
				fmt.Println("Create confgi file")
				os.Exit(0)
			}
		}

		if params.documentNumber == "" {
			return rootCmd.Help()
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
			os.Exit(1)
		}

		return nil
	}

}
