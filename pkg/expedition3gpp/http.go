package expedition3gpp

import (
	"os"
	"io"
	"errors"
	"strings"
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

func setArchiveUrl(url string) archiveUrl {
	a := archiveUrl{url: url}
	return a
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
func getHTMLContents(url string, c chan []Specification) {
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
	// return spec
	c <- spec
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
