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

// --------------------------------------------------
// Search command method
// --------------------------------------------------
func SearchExpedition3gpp(config *Config) error {
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
		| tpppYaml.validateLocation | false |
		+---------------------------+-------+
	*/
	if config.DocumentNumber != "" && !(tpppYaml.validateLocation()) {
		srcUrl := createUrl(config.DocumentNumber)
		spec := getHTMLContents(srcUrl)

		if config.DocumentVersion == "" {
			formatOutput(spec)
		}

		if config.DocumentVersion != "" {
			formatOutputOneVersion(spec, config.DocumentVersion)
		}

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
		| tpppYaml.validateLocation | true  |
		+---------------------------+-------+
	*/
	if config.DocumentNumber != "" && tpppYaml.validateLocation() {
		cy := getCacheValue(config.DocumentNumber)

		if config.DocumentVersion == "" {
			formatOutputYaml(cy)
		}

		if config.DocumentVersion != "" {
			formatOutputYamlOneVersion(cy, config.DocumentVersion)
		}
		return nil
	}
	return errors.New("Assume no error here.")
}

// --------------------------------------------------
// Download command method
// --------------------------------------------------
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

	// cp := getConfigParameter()
	tpppYaml := saveLocation{path: getHomedir() + getSeparate() + notationAdjustment(config.DocumentNumber) + ".yaml"}

	/*
		+---------------------------+-------+
		| Parameter                 | value |
		+---------------------------+-------+
		| config.DocumentNumber     | xxxxx |
		| tpppYaml.validateLocation | false |
		+---------------------------+-------+
	*/
	if config.DocumentNumber != "" && config.DocumentVersion != "" && !(tpppYaml.validateLocation()) {
		srcUrl := createUrl(config.DocumentNumber)
		spec := getHTMLContents(srcUrl)
		var dstUrl *string

		for i, _ := range spec {
			if spec[i].version == config.DocumentVersion {
				dstUrl = &spec[i].url
				break
			}

			if i + 1 == len(spec) {
				return errors.New("The relevant version does not exist.")
			}
		}

		reString := config.DocumentNumber + `-.*zip`
		re := regexp.MustCompile(reString)
		searchResult := re.FindAllStringSubmatch(*dstUrl, -1)

		if len(searchResult) == 0 {
			return errors.New("searchResult is empty")
		}

		filePath := saveLocation{path: getHomedir() + getSeparate() + searchResult[0][0]}

		if filePath.validateLocation() {
			return errors.New("The path specified is not correct.")
		}

		a := archiveUrl{url: *dstUrl}
		if err := a.downloadDocument(filePath.path); err != nil {
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
