package expedition3gpp

import (
	"io"
	"os"
	"fmt"
    "time"
    "strings"
	"archive/zip"
	"path/filepath"
)

// --------------------------------------------------
// File Remove
// --------------------------------------------------
func fileRemove(filePath string) error {
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

// --------------------------------------------------
// File Un zip
// --------------------------------------------------
func fileUnzip(filePath string) error {
    r, err := zip.OpenReader(filePath)
    if err != nil {
        return err
    }
    defer r.Close()

	for _, f := range r.File {
        rc, err := f.Open()
        if err != nil {
            return err
        }
        defer rc.Close()

        path := filepath.Join(getSeparate(), f.Name)
        if f.FileInfo().IsDir() {
            os.MkdirAll(path, f.Mode())

		} else {
            f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
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
    permissionChange(filePath)
    return nil
}

// --------------------------------------------------
// Output
// --------------------------------------------------
func formatOutput(spec []Specification) {
	fmt.Println("+-----+---------+----------------------------------------------------------------------------------+")
	fmt.Println("| No. | Version | URL                                                                              |")
	fmt.Println("+-----+---------+----------------------------------------------------------------------------------+")
	for i := 0; i < len(spec); i++ {
		fmt.Printf("| %3d | %7s | %-80s |\n", i + 1, spec[i].version, spec[i].url)
	}
	fmt.Println("+-----+---------+----------------------------------------------------------------------------------+")
}

func formatOutputOneVersion(spec []Specification, version string) {
	fmt.Println("+-----+---------+----------------------------------------------------------------------------------+")
	fmt.Println("| No. | Version | URL                                                                              |")
	fmt.Println("+-----+---------+----------------------------------------------------------------------------------+")
	for i := 0; i < len(spec); i++ {
        if spec[i].version == version {
            fmt.Printf("| %3d | %7s | %-80s |\n", 1, spec[i].version, spec[i].url)
        }
	}
	fmt.Println("+-----+---------+----------------------------------------------------------------------------------+")
}

func formatOutputYaml(cy cacheYaml) {
	fmt.Println("+-----+---------+----------------------------------------------------------------------------------+")
	fmt.Println("| No. | Version | URL                                                                              |")
	fmt.Println("+-----+---------+----------------------------------------------------------------------------------+")
	for i := 0; i < len(cy.Value); i++ {
		fmt.Printf("| %3d | %7s | %-80s |\n", i + 1, cy.Value[i].Version, cy.Value[i].Url)
	}
	fmt.Println("+-----+---------+----------------------------------------------------------------------------------+")
}

func formatOutputYamlOneVersion(cy cacheYaml, version string) {
	fmt.Println("+-----+---------+----------------------------------------------------------------------------------+")
	fmt.Println("| No. | Version | URL                                                                              |")
	fmt.Println("+-----+---------+----------------------------------------------------------------------------------+")
	for i := 0; i < len(cy.Value); i++ {
        if cy.Value[i].Version == version {
            fmt.Printf("| %3d | %7s | %-80s |\n", 1, cy.Value[i].Version, cy.Value[i].Url)
        }
	}
	fmt.Println("+-----+---------+----------------------------------------------------------------------------------+")
}

// --------------------------------------------------
// Create Cache file
// --------------------------------------------------
type cache struct {
	YamlVersion    int
	Title          string
	CreateDate     string
	Value          []value
}

type value struct {
	Version string
	Name    string
	Url     string
}

func createCacheFile(docNum string, spec []Specification) {
    cache := cache{
        YamlVersion: 2,
        Title:       "3GPP Document " + docNum,
        CreateDate:  getNowTime(),
        Value:       valueStructCreation(docNum, spec),
    }
    fp, b := checkYaml(docNum)
    if b {
        cache.createYaml(fp)
    }
}

func getNowTime() string {
	t := time.Now()
	const layout = "2006-01-02 15:04:05.757"
	return t.Format(layout)
}

func valueStructCreation(docNum string, spec []Specification) []value {
    values := []value{}
    for _, v := range spec {
        value := value{
            Version: v.version,
            Name:    docNum,
            Url:     v.url,
        }
        values = append(values, value)
    }
    return values
}
