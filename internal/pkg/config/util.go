package config

import "github.com/monandkey/expedition3gpp/internal/pkg/fileutil"

// fileCreate is a function to call fileutil's FileCreate.
func fileCreate(fileName string) error {
	return fileutil.FileCreate(fileName)
}

// fileOpen is a function to call fileutil's FileOpen.
func fileOpen(fileName string) error {
	return fileutil.FileOpen(fileName)
}

// fileExist is a function to call fileutil's FileExist.
func fileExist(fileName string) bool {
	return fileutil.FileExist(fileName)
}

// setHomedir is a function to call fileutil's GetHomedir.
func setHomedir() string {
	return fileutil.GetHomedir()
}

// setSeparate is a function to call fileutil's GetSeparate.
func setSeparate() string {
	return fileutil.GetSeparate()
}
