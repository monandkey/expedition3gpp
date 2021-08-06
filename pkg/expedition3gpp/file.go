package expedition3gpp

import (
	"io"
	"os"
	"fmt"
	"log"
	"time"
	"archive/zip"
	"path/filepath"
	"gopkg.in/yaml.v2"
)

// --------------------------------------------------
// File Struct
// --------------------------------------------------
// /home/testusr/xx.xxxx.zip
type saveLocation struct {
	path string
}

// --------------------------------------------------
// File Exist check
// --------------------------------------------------
func (s saveLocation) validateLocation() bool {
	_, err := os.Stat(s.path)
	return err == nil
}

// --------------------------------------------------
// File Remove
// --------------------------------------------------
func (s saveLocation) fileRemove() error {
	if err := os.Remove(s.path); err != nil {
		return err
	}
	return nil
}

// --------------------------------------------------
// File Un zip
// --------------------------------------------------
func (s saveLocation) fileUnzip() error {
    r, err := zip.OpenReader(s.path)
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

        path := filepath.Join(getHomedir() + getSeparate(), f.Name)
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
// Cache struct
// --------------------------------------------------
type cacheYaml struct {
	YamlVersion    int          `yaml:"version"`
	Title          string       `yaml:"title"`
	CreateDate     string       `yaml:"createDate"`
	Value          []valueYaml  `yaml:"value"`
}

type valueYaml struct {
	Version string `yaml:"version"`
	Name    string `yaml:"name"`
	Url     string `yaml:"url"`
}

type cacheFile struct {
	name string
}

// --------------------------------------------------
// Load cache
// --------------------------------------------------
func (c cacheFile) yamlLoad() cacheYaml {
	cacheYaml := cacheYaml{}
	b, _ := os.ReadFile(c.name)
	yaml.Unmarshal(b, &cacheYaml)
	return cacheYaml
}

func (c cacheFile) validateLocation() bool {
	_, err := os.Stat(c.name)
	return err == nil
}

func getCacheValue(d string) cacheYaml {
	fp := getHomedir() + getSeparate() + d + ".yaml"
	cf := cacheFile{name: fp}
	if !(cf.validateLocation()) {
		os.Exit(0)
	}
	return cf.yamlLoad()
}

// --------------------------------------------------
// Create cache
// --------------------------------------------------
func (c cache) createYaml(fp string) {
	f, err := os.OpenFile(fp, os.O_WRONLY|os.O_CREATE, 0664)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	d := yaml.NewEncoder(f)
	if err := d.Encode(&c); err != nil {
		log.Fatal(err)
	}
	defer d.Close()
}

func checkYaml(d string) (string, bool) {
	fp := getHomedir() + getSeparate() + d + ".yaml"
	cf := cacheFile{name: fp}
	if cf.validateLocation() {
		return "", false
	}
	return fp ,true
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
        Title:       "\"3GPP Document " + docNum + "\"",
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
