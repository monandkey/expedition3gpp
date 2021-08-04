package expedition3gpp

import (
	"os"
	"log"
	"fmt"
	"regexp"
)

type Config struct {
	DocumentNumber  string
	DocumentVersion string
	OutputPath      string
	Cache           bool
}

func showUrlList_useNumber(config *Config) error {
	srcUrl := CreateUrl(config.DocumentNumber)
	GetPage(srcUrl)
	return nil
}

func getUrlContnts(config *Config) error {
	srcUrl := CreateUrl(config.DocumentNumber)

	if dstUrl := GetDstUrl(srcUrl, config.DocumentVersion); dstUrl == "0" {
		fmt.Println("The relevant version does not exist.")

	} else {
		reString := config.DocumentNumber + `-.*zip`
		re := regexp.MustCompile(reString)
		searchResult := re.FindAllStringSubmatch(dstUrl, -1)
		filePath := saveLocation{path: getSeparate() + searchResult[0][0]}

		if !(filePath.validateLocation()) {
			os.Exit(0)
		}

		a := archiveUrl{url: dstUrl}
		if err := a.downloadDocument(filePath); err != nil {
			log.Fatal(err)
			os.Exit(0)

		} else {
			if err := FileUnzip(filePath.path); err != nil {
				log.Fatal(err)

			} else {
				if err := FileRemove(filePath.path); err != nil {
					log.Fatal(err)
				}
			}
		}
	}
	return nil
}

func RunExpedition3gpp(config *Config) error {
	if config.DocumentNumber != "" && config.DocumentVersion == "" {
		showUrlList_useNumber(config)

	} else if config.DocumentNumber != "" && config.DocumentVersion != "" {
		getUrlContnts(config)
	}
	return nil
}