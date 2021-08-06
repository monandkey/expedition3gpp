package expedition3gpp

import (
	"errors"
	"regexp"
)

type Config struct {
	DocumentNumber  string
	DocumentVersion string
	OutputPath      string
	Cache           bool
}

func RunExpedition3gpp(config *Config) error {
	// Assumption to be executed only the first time.
	// Runs if the config file does not exist.
	if ExistInitConfig() {
		initConfig := InitConfig{
			StrageLocation:     "HOMEDIR",
			CacheEnable:        true,
			CacheRetentionTime: 14400,
			CacheLocation:      "HOMEDIR",
		}
		InitializeConfig(&initConfig)
	}

	cp := getConfigParameter()
	tpppYaml := saveLocation{path: getHomedir() + getSeparate() + notationAdjustment(config.DocumentNumber) + ".yaml"}

	/*
		+---------------------------+-------+
		| Parameter                 | value |
		+---------------------------+-------+
		| config.DocumentNumber     | xxxxx |
		| config.DocumentVersion    | ""    |
		| tpppYaml.validateLocation | false |
		+---------------------------+-------+
	*/
	if config.DocumentNumber != "" && config.DocumentVersion == "" && !(tpppYaml.validateLocation()) {
		srcUrl := createUrl(config.DocumentNumber)
		spec := getHTMLContents(srcUrl)
		formatOutput(spec)

		if cp.CacheEnable {
			createCacheFile(config.DocumentNumber, spec)
		}
		return nil
	}

	/*
		+---------------------------+-------+
		| Parameter                 | value |
		+---------------------------+-------+
		| config.DocumentNumber     | xxxxx |
		| config.DocumentVersion    | ""    |
		| tpppYaml.validateLocation | true  |
		+---------------------------+-------+
	*/
	if config.DocumentNumber != "" && config.DocumentVersion == "" && tpppYaml.validateLocation() {
		cy := getCacheValue(config.DocumentNumber)
		formatOutputYaml(cy)
		return nil
	}

	/*
		+---------------------------+-------+
		| Parameter                 | value |
		+---------------------------+-------+
		| config.DocumentNumber     | xxxxx |
		| config.DocumentVersion    | x.x.x |
		| tpppYaml.validateLocation | false |
		+---------------------------+-------+
	*/
	if config.DocumentNumber != "" && config.DocumentVersion != "" && !(tpppYaml.validateLocation()) {
		srcUrl := createUrl(config.DocumentNumber)
		spec := getHTMLContents(srcUrl)
		formatOutputOneVersion(spec, config.DocumentVersion)

		if cp.CacheEnable {
			createCacheFile(config.DocumentNumber, spec)
		}
		return nil
	}

	/*
		+---------------------------+-------+
		| Parameter                 | value |
		+---------------------------+-------+
		| config.DocumentNumber     | xxxxx |
		| config.DocumentVersion    | x.x.x |
		| tpppYaml.validateLocation | true  |
		+---------------------------+-------+
	*/
	if config.DocumentNumber != "" && config.DocumentVersion != "" && tpppYaml.validateLocation() {
		cy := getCacheValue(config.DocumentNumber)
		formatOutputYamlOneVersion(cy, config.DocumentVersion)
		return nil
	}

	if config.DocumentNumber != "" && config.DocumentVersion != "" && !(tpppYaml.validateLocation()) {
		srcUrl := createUrl(config.DocumentNumber)
		dstUrl := getDstUrl(srcUrl, config.DocumentVersion);
	
		if dstUrl == "" {
			return errors.New("The relevant version does not exist.")
		}
	
		reString := config.DocumentNumber + `-.*zip`
		re := regexp.MustCompile(reString)
		searchResult := re.FindAllStringSubmatch(dstUrl, -1)
		filePath := saveLocation{path: getSeparate() + searchResult[0][0]}
	
		if !(filePath.validateLocation()) {
			return errors.New("The path specified is not correct.")
		}
	
		a := archiveUrl{url: dstUrl}
		if err := a.downloadDocument(filePath); err != nil {
			return err
		}
	
		if err := fileUnzip(filePath.path); err != nil {
			return err
		}
	
		if err := fileRemove(filePath.path); err != nil {
			return err
		}
	}
	return nil
}
