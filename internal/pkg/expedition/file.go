package expedition

import (
	"io/ioutil"
	"regexp"

	"gopkg.in/yaml.v2"
)

func getSaveFilePath(outputPath string, documentNumber string, url string) string {
	if outputPath == "HOMEDIR" {
		outputPath = getHomedir()
	}
	reString := documentNumber + `-.*zip`
	re := regexp.MustCompile(reString)
	searchResult := re.FindAllStringSubmatch(url, -1)

	outputPath = regexp.MustCompile(`/$`).ReplaceAllString(outputPath, "")
	path := outputPath + "/" + searchResult[0][0]
	return path
}

func getSaveCachePath(configParams configParams, documentNumber string) string {
	var outputPath string
	if configParams.cacheLocation == "HOMEDIR" {
		outputPath = getHomedir()
	}

	outputPath = regexp.MustCompile(`/$`).ReplaceAllString(outputPath, "")
	path := outputPath + "/.cache/" + documentNumber + ".yaml"
	return path
}

func yamlWrite(filepath string, data interface{}) error {
	if fileExist(filepath) {
		if err := fileOpen(filepath); err != nil {
			return err
		}
	} else {
		_, err := fileCreate(filepath)
		if err != nil {
			return nil
		}
	}

	buf, err := yaml.Marshal(data)
	if err != nil {
		return err
	}

	if err := ioutil.WriteFile(filepath, buf, 0664); err != nil {
		return err
	}
	return nil
}
