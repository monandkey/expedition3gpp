package expedition

import (
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/monandkey/expedition3gpp/internal/pkg/config"
	"gopkg.in/yaml.v2"
)

// configLoad is a function that transfers the loaded configuration to another structure.
func configLoad() configParams {
	c := config.SelectConfigUser()
	value := c.Load()
	configParams := configParams{
		strageLocation:     value.StrageLocation,
		cacheEnable:        value.CacheEnable,
		cacheRetentionTime: value.CacheRetentionTime,
		cacheLocation:      value.CacheLocation,
	}
	return configParams
}

// getCacheFileName is a function to get the name of the cache file.
func getCacheFileName(cacheLocation string, documentNumber string) string {
	var filePath string
	docNum := notationAdjustment(documentNumber)
	separate := getSeparate()
	re := regexp.MustCompile(separate + `$`)

	if cacheLocation == "HOMEDIR" {
		filePath = getHomedir() + separate + ".cache" + separate + docNum + ".yaml"

	} else if re.MatchString(cacheLocation) {
		filePath = cacheLocation + ".cache" + separate + docNum + ".yaml"

	} else if !(re.MatchString(cacheLocation)) {
		filePath = cacheLocation + separate + ".cache" + separate + docNum + ".yaml"
	}
	return filePath
}

// notationAdjustment function is used to remove a dot if there is one.
func notationAdjustment(documentNumber string) string {
	if strings.Contains(documentNumber, ".") {
		return strings.Replace(documentNumber, ".", "", 1)
	}
	return documentNumber
}

// cacheValidate is a function to validate that the loaded cache is still valid.
func cacheValidate(cacheRetentionTime int, fileName string) bool {
	if !(fileExist(fileName)) {
		return false
	}

	yamlStruct := yamlLoad(fileName)
	layout := "2006-01-02 15:04:05"

	/*
		+--------------------------------+
		| name | description             |
		+--------------------------------+
		| t1   | YAML input Date         |
		| t2   | t1 + CacheRetentionTime |
		| t3   | Current date            |
		+--------------------------------+
	*/
	t1, _ := time.Parse(layout, yamlStruct.CreateDate)
	t2 := t1.AddDate(0, 0, cacheRetentionTime/1440)
	t3 := time.Now()

	/*
		+------------+------------+-------+
		| t2         | t3         | bool  |
		+------------+------------+-------+
		| 2021-08-10 | 2021-08-11 | False |
		| 2021-08-10 | 2021-08-09 | True  |
		+------------+------------+-------+
	*/
	return t3.Before(t2)
}

// yamlLoad is a function to load a cache file.
func yamlLoad(fileName string) yamlStruct {
	yamlStruct := yamlStruct{}
	b, _ := os.ReadFile(fileName)
	yaml.Unmarshal(b, &yamlStruct)
	return yamlStruct
}

// cacheLoad is a function to call yamlLoad.
func cacheLoad(filePath string) yamlStruct {
	return yamlLoad(filePath)
}
