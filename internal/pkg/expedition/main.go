/*
This is the main package of expedition.
*/
package expedition

import (
	"errors"
	"fmt"
	"regexp"
)

// SelectUser is a function that returns an interface.
func SelectUser() ExpeditionAction {
	return &baseParams{}
}

// SetParams is a function to set the required parameters.
func (b *baseParams) SetParams(
	documentNumber string,
	documentVersion string,
	outputPath string,
	cache bool,
	releaseNumber string,
) {
	b.params.DocumentNumber = documentNumber
	b.params.DocumentVersion = documentVersion
	b.params.OutputPath = outputPath
	b.params.Cache = cache
	b.params.releaseNumber = releaseNumber
}

// Search is a function to search for 3GPP documents.
func (b *baseParams) Search() error {
	config := configLoad()
	b.DocumentNumber = notationAdjustment(b.DocumentNumber)
	filePath := getCacheFileName(config.cacheLocation, b.DocumentNumber)

	var err error
	if cacheValidate(config.cacheRetentionTime, filePath) && !(b.params.Cache) {
		contents := cacheLoad(filePath)
		b.value = contents.Value
	} else {
		cancel := make(chan struct{})
		go displayLoading(cancel)
		b.value, err = pageFetch(b.DocumentNumber)
		close(cancel)

		if err != nil {
			return err
		}
		fmt.Printf("\r[OK] Download Success.\n")
	}

	if b.releaseNumber != "" {
		formatDisplayRelease(b.value, b.releaseNumber)
	} else if b.DocumentVersion != "" {
		formatDisplayVersion(b.value, b.DocumentVersion)
	} else {
		formatDisplayAll(b.value)
	}
	return nil
}

// Download is a function to download the 3GPP document.
func (b *baseParams) Download() error {
	if b.DocumentVersion == "" {
		fmt.Printf("Please specify the version you want to download. : ")
		fmt.Scan(&b.DocumentVersion)
	}

	downloadUrl := getDownloadURL(b.value, b.DocumentVersion)
	if downloadUrl == "" {
		return errors.New("the specified version does not exist")
	}

	configParams := configLoad()
	if b.OutputPath == "" {
		b.OutputPath = configParams.strageLocation
	}

	filePath := getSaveFilePath(b.OutputPath, b.DocumentNumber, downloadUrl)
	cancel := make(chan struct{})
	go displayLoading(cancel)
	if err := downloadContents(downloadUrl, filePath); err != nil {
		return err
	}
	close(cancel)

	fmt.Printf("\r[OK] Download Success.\n")
	if err := fileUnzip(filePath); err != nil {
		return err
	}

	if err := fileRemove(filePath); err != nil {
		return err
	}

	return nil
}

// Cache is a function for storing the information retrieved by the search in cache.
func (b *baseParams) Cache() error {
	configParams := configLoad()
	if !(configParams.cacheEnable) {
		return nil
	}
	cacheName := getSaveCachePath(configParams, b.DocumentNumber)
	docNum := regexp.MustCompile(`(^.{2})`).ReplaceAllString(b.DocumentNumber, "$1.")
	data := yamlStruct{
		YamlVersion: 2,
		Title:       "3GPP Document " + docNum,
		CreateDate:  getNowTime(),
		Value:       b.value,
	}
	if err := yamlWrite(cacheName, data); err != nil {
		return err
	}
	return nil
}
