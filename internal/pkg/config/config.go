package config

import (
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

func getFileName() string {
	const configFileName string = ".expedition3gpp.yaml"
	homeDir := setHomedir()
	separate := setSeparate()
	return homeDir + separate + configFileName
}

func configLoad(fileName string) params {
	params := params{}
	b, _ := os.ReadFile(fileName)
	yaml.Unmarshal(b, &params)
	return params
}

func configWrite(fileName string, data interface{}) error {
	if fileExist(fileName) {
		if err := fileOpen(fileName); err != nil {
			return err
		}
	} else {
		if err := fileCreate(fileName); err != nil {
			return nil
		}
	}

	buf, err := yaml.Marshal(data)
	if err != nil {
		return err
	}

	if err := ioutil.WriteFile(fileName, buf, 0664); err != nil {
		return err
	}
	return nil
}
