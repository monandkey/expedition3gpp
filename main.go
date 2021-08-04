package main

import (
	"os"
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"local.packages/expedition3gpp"
)

type params struct {
	documentNumber string
	documentVersion string
	outputPath string
	cache bool
}

// --------------------------------------------------
// Main
// --------------------------------------------------
func main() {
	params := params{
		documentNumber:  "",
		documentVersion: "",	
		outputPath:      "current",
		cache:           false,
	}

	initConfig := expedition3gpp.GetConfigParameter()

	cmd := &cobra.Command{}
	cmd.Use = "expedition3gpp"
	cmd.Short = "Download the 3GPP document"

	cmd.Flags().StringVarP(&params.documentNumber, "document-number", "N", params.documentNumber, "3GPP Document Number")
	cmd.Flags().StringVarP(&params.documentVersion, "document-version", "V", params.documentVersion, "3GPP Document Version")
	cmd.Flags().StringVarP(&params.outputPath, "output-path", "O", params.outputPath, "Where you want the output to go")
	cmd.Flags().BoolVar(&params.cache, "no-cache", params.cache, "Not using cache")

	cmd.Flags().Bool("init", true, "Initialize expedition3gpp config")
	cmd.Flags().StringVar(&initConfig.StrageLocation, "strage-location", initConfig.StrageLocation, "")
	cmd.Flags().BoolVar(&initConfig.CacheEnable, "cache-enable", initConfig.CacheEnable, "")
	cmd.Flags().IntVar(&initConfig.CacheRetentionTime, "cache-retention-time", initConfig.CacheRetentionTime, "")
	cmd.Flags().StringVar(&initConfig.CacheLocation, "cache-location", initConfig.CacheLocation, "")

	cmd.RunE = func(cmd *cobra.Command, args []string) error {
		for _, v := range os.Args {
			if strings.Contains(v, "init") || expedition3gpp.ExistInitConfig() {
				expedition3gpp.InitializeConfig()
				fmt.Println("Create confgi file")
				os.Exit(0)
			}
		}

		if params.documentNumber == "" {
			return cmd.Help()
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

	if err := cmd.Execute(); err != nil {
		fmt.Println(os.Stderr, err)
		os.Exit(1)
	}
}