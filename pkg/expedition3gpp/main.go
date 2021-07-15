package expedition3gpp

import (
	"log"
	"fmt"
	"regexp"
)

type Config struct {
	Url string
	DocumentNumber string
	DocumentVersion string
	OutputPath string
}

func showUrlList_useUrl(config *Config) error {
	GetPage(config.Url)
	return nil
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
		filePath := "./" + searchResult[0][0]

		if err := TargetDownload(filePath, dstUrl); err != nil {
			panic(err)

		} else {
			if err := FileUnzip(filePath); err != nil {
				log.Fatal(err)

			} else {
				if err := FileRemove(filePath); err != nil {
					log.Fatal(err)
				}
			}
		}
	}
	return nil
}

func RunExpedition3gpp(config *Config) error {
	if config.Url != "default" {
		showUrlList_useUrl(config)

	} else if config.DocumentNumber != "default" && config.DocumentVersion == "default" {
		showUrlList_useNumber(config)

	} else if config.DocumentNumber != "default" && config.DocumentVersion != "default" {
		getUrlContnts(config)
	}
	return nil
}