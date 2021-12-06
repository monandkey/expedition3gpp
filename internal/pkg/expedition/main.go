package expedition

import (
	"errors"
	"regexp"
)

func SelectUser() ExpeditionAction {
	return &baseParams{}
}

func (b *baseParams) SetParams(
	documentNumber string,
	documentVersion string,
	outputPath string,
	cache bool,
) {
	b.params.DocumentNumber = documentNumber
	b.params.DocumentVersion = documentVersion
	b.params.OutputPath = outputPath
	b.params.Cache = cache
}

func (b *baseParams) Search() error {
	config := configLoad()
	b.DocumentNumber = notationAdjustment(b.DocumentNumber)
	filePath := getCacheFileName(config.cacheLocation, b.DocumentNumber)

	var err error
	if cacheValidate(config.cacheRetentionTime, filePath) {
		contents := cacheLoad(filePath)
		b.value = contents.Value
	} else {
		b.value, err = pageFetch(b.DocumentNumber)
		if err != nil {
			return err
		}
	}
	formatDisplay(b.value)
	return nil
}

func (b *baseParams) Download() error {
	downloadUrl := getDownloadURL(b.value, b.DocumentVersion)
	if downloadUrl == "" {
		return errors.New("the specified version does not exist")
	}

	filePath := getSaveFilePath(b.OutputPath, b.DocumentNumber, downloadUrl)
	err := downloadContents(downloadUrl, filePath)
	if err != nil {
		return err
	}
	return nil
}

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
