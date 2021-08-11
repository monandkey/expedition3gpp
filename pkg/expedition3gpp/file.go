package expedition3gpp

import (
	"io"
	"os"
	"log"
	"fmt"
	"time"
	"regexp"
	"strings"
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

func setSaveLocation(path string) saveLocation {
	s := saveLocation{path: path}
	return s
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

		rep := regexp.MustCompile(`[0-9]{5}-...\.zip`)
		path := filepath.Join(rep.ReplaceAllString(s.path, f.Name))

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
// Strage Location
// --------------------------------------------------
func strageLocation(d string) string {
	cp := getConfigParameter()
	if cp.StrageLocation == "HOMEDIR" {
		s := getHomedir() + getSeparate() + d
		return s
	
	} else if cp.StrageLocation != "HOMEDIR" && strings.HasSuffix(cp.StrageLocation, getSeparate()) {
		s := cp.StrageLocation + d
		return s

	} else if cp.StrageLocation != "HOMEDIR" && !(strings.HasSuffix(cp.StrageLocation, getSeparate())) {
        s := cp.StrageLocation + getSeparate() + d
		return s

	}
	fmt.Println("The specified path does not exist.")
	os.Exit(0)
	return ""
}

func outputLocation(o string, d string) string {
	if strings.HasSuffix(o, getSeparate()) {
		s := o + d
		return s
	
	} else if !(strings.HasSuffix(o, getSeparate())) {
		s := o + getSeparate() + d
		return s
	}
	fmt.Println("The specified path does not exist.")
	os.Exit(0)
	return ""
}

// --------------------------------------------------
// Cache struct
// --------------------------------------------------
type cacheYaml struct {
	YamlVersion    int          `yaml:"version"`
	Title          string       `yaml:"title"`
	CreateDate     string       `yaml:"createdate"`
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
	fp := cacheLocation(d)
	cf := cacheFile{name: fp}
	if !(cf.validateLocation()) {
		os.Exit(0)
	}
	return cf.yamlLoad()
}

func cacheTimeVerification(ct string, ci int) bool {
	layout := "2006-01-02 15:04:05"

	/*
		+--------------------------------+
		| name | description             |
		+--------------------------------+
		| t1   | YAML input Date         |
		| t2   | t1 + CacheRetentionTime |
		| t3   | Current date            |
		+--------------------------------+
	*/
	t1, _ := time.Parse(layout, ct)
	t2 := t1.AddDate(0, 0, ci/1440)
	t3 := time.Now()

	/*
		+------------+------------+-------+
		| t2         | t3         | bool  |
		+------------+------------+-------+
		| 2021-08-10 | 2021-08-11 | False |
		| 2021-08-10 | 2021-08-09 | True  |
		+------------+------------+-------+
	*/
	if t3.Before(t2) {
		return false
	}
	return true
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
	fp := cacheLocation(d)
	cf := cacheFile{name: fp}
	if cf.validateLocation() {
		return "", false
	}
	return fp ,true
}

func cacheLocation(d string) string {
	cp := getConfigParameter()
	if cp.CacheLocation == "HOMEDIR" {
		s := getHomedir() + getSeparate() + ".cache" + getSeparate() + d + ".yaml"
		return s
	
	} else if cp.CacheLocation != "HOMEDIR" && strings.Contains(cp.CacheLocation, getSeparate()) {
		s := cp.CacheLocation + ".cache" + getSeparate() + d + ".yaml"
		return s

	} else if cp.CacheLocation != "HOMEDIR" && !(strings.Contains(cp.CacheLocation, getSeparate())) {
		s := cp.CacheLocation + getSeparate() + ".cache" + getSeparate() + d + ".yaml"
		return s

	}
	fmt.Println("The specified path does not exist.")
	os.Exit(0)
	return ""
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
