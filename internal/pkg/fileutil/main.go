package fileutil

import (
	"os"
	"runtime"
)

func FileExist(fileName string) bool {
	_, err := os.Stat(fileName)
	return err == nil
}

func FileCreate(fileName string) error {
	fp, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer fp.Close()
	return nil
}

func FileOpen(fileName string) error {
	f, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE, 0664)
	if err != nil {
		return err
	}
	defer f.Close()
	return nil
}

func GetHomedir() string {
	h, err := os.UserHomeDir()
	if err != nil {
		return ""
	}
	return h
}

func GetSeparate() string {
	switch runtime.GOOS {
	case "windows":
		return "\\"
	case "linux":
		return "/"
	}
	return ""
}
