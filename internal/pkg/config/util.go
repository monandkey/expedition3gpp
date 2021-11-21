package config

import "github.com/monandkey/expedition3gpp/internal/pkg/fileutil"

func fileCreate(fileName string) error {
	return fileutil.FileCreate(fileName)
}

func fileOpen(fileName string) error {
	return fileutil.FileOpen(fileName)
}

func fileExist(fileName string) bool {
	return fileutil.FileExist(fileName)
}

func setHomedir() string {
	return fileutil.GetHomedir()
}

func setSeparate() string {
	return fileutil.GetSeparate()
}
