package expedition

import (
	"os"

	"github.com/monandkey/expedition3gpp/internal/pkg/fileutil"
)

// getHomedir is a function to call fileutil's GetHomedir.
func getHomedir() string {
	return fileutil.GetHomedir()
}

// getSeparate is a function to call fileutil's GetSeparate.
func getSeparate() string {
	return fileutil.GetSeparate()
}

// fileExist is a function to call fileutil's FileExist.
func fileExist(fileName string) bool {
	return fileutil.FileExist(fileName)
}

// fileCreate is a function to call fileutil's FileCreateReturnAll.
func fileCreate(fileName string) (*os.File, error) {
	out, err := fileutil.FileCreateReturnAll(fileName)
	return out, err
}

// fileRemove is a function to call fileutil's FileRemove.
func fileRemove(fileName string) error {
	return fileutil.FileRemove(fileName)
}

// fileOpen is a function to call fileutil's FileOpen.
func fileOpen(fileName string) error {
	return fileutil.FileOpen(fileName)
}
