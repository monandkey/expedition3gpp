package expedition3gpp

import (
	"os"
	"io"
	"errors"
	"regexp"
	"net/http"

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
	return os.IsNotExist(err)
}

func (a archiveUrl) downloadDocument(f saveLocation) error {
	resp, err := http.Get(a.url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	out, err := os.Create(f.path)
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
func GetPage(url string) {
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
	formatOutput(spec)
}

func GetDstUrl(url string, docVer string) string {
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
	for i := 0; i < len(spec); i++ {
		if spec[i].version == docVer {
			return spec[i].url
		} else {
			continue
		}
	}
	return "0"
}

// --------------------------------------------------
// Create URL
// --------------------------------------------------
func CreateUrl(docNum string) string {
	srcUrl := "https://www.3gpp.org/DynaReport/" + docNum + ".htm"
	return srcUrl
}

