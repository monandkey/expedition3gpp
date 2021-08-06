package expedition3gpp

import (
	"os"
	"io"
	"log"
	"errors"
	"strings"
	"regexp"
	"net/http"

	"gopkg.in/yaml.v2"
	"github.com/PuerkitoBio/goquery"
)

// --------------------------------------------------
// File Download
// --------------------------------------------------
type archiveUrl struct {
	url string
}

type saveLocation struct {
	path string
}

func (s saveLocation) validateLocation() bool {
	_, err := os.Stat(s.path)
	return err == nil
}

func (a archiveUrl) downloadDocument(f string) error {
	resp, err := http.Get(a.url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	out, err := os.Create(f)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	return err
}

// --------------------------------------------------
// String Search
// --------------------------------------------------
func stringSearch(targetString string, reString string) (string, error) {
	re := regexp.MustCompile(reString)
	searchResult := re.FindAllStringSubmatch(targetString, -1)
	if searchResult != nil {
		return searchResult[0][0], errors.New("param is empty")
	}
	return "0", nil
}

type Specification struct {
	url string
	version string
}

// --------------------------------------------------
// Get HTML
// --------------------------------------------------
func getHTMLContents(url string) []Specification {
	spec := make([]Specification, 0)

	doc, _ := goquery.NewDocument(url)
	doc.Find("a").Each(func(_ int, s *goquery.Selection) {
		href, _ := s.Attr("href")
		text := s.Text()

		str1, err1 := stringSearch(href, `http.*.zip`)
		str2, err2 := stringSearch(text, `[0-9]{1,2}\.[0-9]{1,2}\.[0-9]{1,2}`)
		if err1 != nil && err2 != nil {
			spec = append(spec, Specification{str1, str2})
		}
	})
	return spec
}

// --------------------------------------------------
// Create URL
// --------------------------------------------------
func createUrl(docNum string) string {
	srcUrl := "https://www.3gpp.org/DynaReport/" + notationAdjustment(docNum) + ".htm"
	return srcUrl
}

func notationAdjustment(docNum string) string {
	if strings.Index(docNum, ".") != -1 {
		return strings.Replace(docNum, ".", "", 1)
	}
	return docNum
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
