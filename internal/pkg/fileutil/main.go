/*
Separate common functions used in file operations into packages.
*/
package fileutil

import (
	"os"
	"runtime"
)

// FileExist is a function that checks for the existence of a file.
func FileExist(fileName string) bool {
	_, err := os.Stat(fileName)
	return err == nil
}

// FileCreate is a function that creates a file.
func FileCreate(fileName string) error {
	fp, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer fp.Close()
	return nil
}

// FileCreateReturnAll is a function that creates a file and returns it including error.
func FileCreateReturnAll(fileName string) (*os.File, error) {
	out, err := os.Create(fileName)
	return out, err
}

// FileRemove is a function to delete a file.
func FileRemove(fileName string) error {
	if err := os.Remove(fileName); err != nil {
		return err
	}
	return nil
}

// FileOpen is a function to open a file.
func FileOpen(fileName string) error {
	f, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE, 0664)
	if err != nil {
		return err
	}
	defer f.Close()
	return nil
}

// GetHomedir is a function to get the home directory.
func GetHomedir() string {
	h, err := os.UserHomeDir()
	if err != nil {
		return ""
	}
	return h
}

// GetSeparate is a function to get the separation of a file path.
func GetSeparate() string {
	switch runtime.GOOS {
	case "windows":
		return "\\"
	case "linux":
		return "/"
	}
	return ""
}
