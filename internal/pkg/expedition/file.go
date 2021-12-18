package expedition

import (
	"archive/zip"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"

	"gopkg.in/yaml.v2"
)

// getSaveFilePath is a function that returns the path to save the document.
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

// getSaveCachePath is a function that returns the path to save the cache
func getSaveCachePath(configParams configParams, documentNumber string) string {
	var outputPath string
	if configParams.cacheLocation == "HOMEDIR" {
		outputPath = getHomedir()
	}

	outputPath = regexp.MustCompile(`/$`).ReplaceAllString(outputPath, "")
	path := outputPath + "/.cache/" + documentNumber + ".yaml"
	return path
}

// yamlWrite is a function for writing yaml
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

// fileUnzip is a function for unzipping a zip file.
func fileUnzip(path string) error {
	r, err := zip.OpenReader(path)
	if err != nil {
		return err
	}
	defer r.Close()

	for _, v := range r.File {
		rc, err := v.Open()
		if err != nil {
			return err
		}
		defer rc.Close()

		reString := regexp.MustCompile(`[0-9]{5}-...\.zip`)
		fullPath := filepath.Join(reString.ReplaceAllString(path, v.Name))

		if v.FileInfo().IsDir() {
			os.MkdirAll(fullPath, v.Mode())
		} else {
			f, err := os.OpenFile(fullPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, v.Mode())
			if err != nil {
				return err
			}
			defer f.Close()

			_, err = io.Copy(f, rc)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
