package config

import (
	"io/ioutil"
	"os"
	"runtime"

	"gopkg.in/yaml.v2"
)

func getFileName() string {
	const configFileName string = ".expedition3gpp.yaml"
	homeDir := setHomedir()
	separate := setSeparate()
	return homeDir + separate + configFileName
}

func setHomedir() string {
	h, err := os.UserHomeDir()
	if err != nil {
		return ""
	}
	return h
}

func setSeparate() string {
	switch runtime.GOOS {
	case "windows":
		return "\\"
	case "linux":
		return "/"
	}
	return ""
}

func fileExist(fileName string) bool {
	_, err := os.Stat(fileName)
	return err == nil
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

func fileCreate(fileName string) error {
	fp, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer fp.Close()
	return nil
}

func fileOpen(fileName string) error {
	f, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE, 0664)
	if err != nil {
		return err
	}
	defer f.Close()
	return nil
}
