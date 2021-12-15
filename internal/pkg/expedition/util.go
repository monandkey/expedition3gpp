package expedition

import (
	"os"

	"github.com/monandkey/expedition3gpp/internal/pkg/fileutil"
)

func getHomedir() string {
	return fileutil.GetHomedir()
}

func getSeparate() string {
	return fileutil.GetSeparate()
}

func fileExist(fileName string) bool {
	return fileutil.FileExist(fileName)
}

func fileCreate(fileName string) (*os.File, error) {
	out, err := fileutil.FileCreateReturnAll(fileName)
	return out, err
}

func fileRemove(fileName string) error {
	return fileutil.FileRemove(fileName)
}

func fileOpen(fileName string) error {
	return fileutil.FileOpen(fileName)
}
