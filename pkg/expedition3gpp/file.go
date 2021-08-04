package expedition3gpp

import (
	"os"
    "strings"
)

// --------------------------------------------------
// File Remove
// --------------------------------------------------
func FileRemove(filePath string) error {
	if err := os.Remove(filePath); err != nil {
		return err
	}
	return nil
}

// --------------------------------------------------
// File PermissionChange change
// --------------------------------------------------
func fileExists(filename string) bool {
    _, err := os.Stat(filename)
    return err == nil
}

func permissionChange(filePath string) {
	srcString := strings.Replace(filePath, "zip", "", 1)

	targetString := srcString + "doc"
	if fileExists(targetString) {
		_ = os.Chmod(targetString, 0600)
	}

	if fileExists(targetString + "x") {
		os.Chmod(targetString + "x", 0600)
	}
}
