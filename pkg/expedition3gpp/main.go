package expedition3gpp

import (
	"errors"
	"fmt"
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
	tpppYaml := setSaveLocation(cacheLocation(notationAdjustment(config.DocumentNumber)))

	if tpppYaml.validateLocation() || config.Cache {
		cy := getCacheValue(config.DocumentNumber)
		if cacheTimeVerification(cy.CreateDate, cp.CacheRetentionTime) || config.Cache {
			err := tpppYaml.fileRemove()
			if err != nil {
				return err
			}
		}
	}

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
		c := make(chan []specDocInfo)
		cancel := make(chan struct{})

		go fetch3gppPage(srcUrl, c)
		go displayLoading(cancel)

		spec := <-c
		close(cancel)

		if len(spec) == 0 {
			return errors.New("\rThe specified document does not exist.     ")
		}
		fmt.Printf("\r[OK] Download Success.\n")

		if config.DocumentVersion == "" {
			formatOutput(spec)

		} else if config.DocumentVersion != "" {
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

		} else if config.DocumentVersion != "" {
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

	cp := getConfigParameter()
	tpppYaml := setSaveLocation(cacheLocation(notationAdjustment(config.DocumentNumber)))
	var dstUrl *string
	cancel := make(chan struct{})

	if tpppYaml.validateLocation() || config.Cache {
		cy := getCacheValue(config.DocumentNumber)
		if cacheTimeVerification(cy.CreateDate, cp.CacheRetentionTime) || config.Cache {
			err := tpppYaml.fileRemove()
			if err != nil {
				return err
			}
		}
	}

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
		c := make(chan []specDocInfo)

		go fetch3gppPage(srcUrl, c)
		go displayLoading(cancel)

		spec := <-c
		close(cancel)

		if len(spec) == 0 {
			return errors.New("\rThe specified document does not exist.     ")
		}
		fmt.Printf("\r[OK] Download Success.\n")

		var verNum string
		if config.DocumentVersion == "" {
			formatOutput(spec)
			fmt.Printf("Please specify the version you want to download. : ")
			fmt.Scan(&verNum)
		}

		for i := range spec {
			if spec[i].version == config.DocumentVersion || spec[i].version == verNum {
				dstUrl = &spec[i].url
				break
			}

			if i+1 == len(spec) {
				return errors.New("The relevant version does not exist.")
			}
		}

		if cp.CacheEnable {
			createCacheFile(config.DocumentNumber, spec)
		}

		/*
			+---------------------------+-------+
			| Parameter                 | value |
			+---------------------------+-------+
			| config.DocumentNumber     | xxxxx |
			| tpppYaml.validateLocation | true  |
			+---------------------------+-------+
		*/
	} else if config.DocumentNumber != "" && tpppYaml.validateLocation() {
		cy := getCacheValue(config.DocumentNumber)

		var verNum string
		if config.DocumentVersion == "" {
			formatOutputYaml(cy)
			fmt.Printf("Please specify the version you want to download. : ")
			fmt.Scan(&verNum)
		}

		for i := range cy.Value {
			if cy.Value[i].Version == config.DocumentVersion || cy.Value[i].Version == verNum {
				dstUrl = &cy.Value[i].Url
				break
			}

			if i+1 == len(cy.Value) {
				return errors.New("The relevant version does not exist.")
			}
		}

	} else {
		return errors.New("Func : RunExpedition3gpp : An unexpected error has occurred.")
	}

	// After
	reString := config.DocumentNumber + `-.*zip`
	re := regexp.MustCompile(reString)
	searchResult := re.FindAllStringSubmatch(*dstUrl, -1)

	if len(searchResult) == 0 {
		return errors.New("searchResult is empty")
	}

	var filePath saveLocation
	if config.OutputPath == "" {
		filePath = setSaveLocation(strageLocation(searchResult[0][0]))

	} else {
		filePath = setSaveLocation(outputLocation(config.OutputPath, searchResult[0][0]))
	}

	if filePath.validateLocation() {
		return errors.New("The path specified is not correct.")
	}

	go displayLoading(cancel)
	a := setArchiveUrl(*dstUrl)
	if err := a.downloadDocument(filePath.path); err != nil {
		return err
	}
	close(cancel)
	fmt.Printf("\r[OK] Download Success.\n")

	if err := filePath.fileUnzip(); err != nil {
		return err
	}

	if err := filePath.fileRemove(); err != nil {
		return err
	}

	fmt.Printf("The files are stored in the following locations\n")
	fmt.Printf("PATH: %s\n", filePath.path)

	return nil
}
