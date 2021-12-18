package expedition

import (
	"io"
	"net/http"
	"regexp"

	"github.com/PuerkitoBio/goquery"
)

// getRequest is a function for issuing a GET Request
func getRequest(url string) (*http.Response, error) {
	res, err := http.Get(url)
	if err != nil {
		return res, err
	}

	if res.StatusCode != 200 {
		return res, getErrorMessage(res.StatusCode)
	}
	return res, nil
}

// pageFetch is a function for retrieving HTML content
func pageFetch(docNum string) ([]valueBody, error) {
	srcUrl := "https://www.3gpp.org/DynaReport/" + docNum + ".htm"
	contents := []valueBody{}
	res, err := getRequest(srcUrl)
	if err != nil {
		return contents, err
	}
	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return contents, err
	}
	contents = pageParse(doc)
	return contents, nil
}

// pageParse is a function for parsing the retrieved HTML content
func pageParse(doc *goquery.Document) []valueBody {
	contents := []valueBody{}
	doc.Find("a").Each(func(i int, s *goquery.Selection) {
		href, _ := s.Attr("href")
		text := s.Text()
		zipUrl := stringSearch(href, `http.*.zip`)
		versionInfo := stringSearch(text, `[0-9]{1,2}\.[0-9]{1,2}\.[0-9]{1,2}`)
		if zipUrl != "" && versionInfo != "" {
			data := valueBody{
				Url:     zipUrl,
				Name:    regexp.MustCompile(`[0-9]{2}.[0-9]{3}`).FindString(zipUrl),
				Version: versionInfo,
			}
			contents = append(contents, data)
		}
	})
	return contents
}

// stringSearch is a function to retrieve the required string from targetString by regular expression
func stringSearch(targetString string, reString string) string {
	re := regexp.MustCompile(reString)
	searchResult := re.FindAllStringSubmatch(targetString, -1)
	if searchResult != nil {
		return searchResult[0][0]
	}
	return ""
}

// downloadContents is a function to download 3GPP documents.
func downloadContents(url string, fileName string) error {
	res, err := getRequest(url)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	out, err := fileCreate(fileName)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, res.Body)
	if err != nil {
		return err
	}
	return nil
}

// getDownloadURL is a function that returns the version number of the document to be downloaded.
func getDownloadURL(valu []valueBody, version string) string {
	for _, v := range valu {
		if v.Version != version {
			continue
		}
		return v.Url
	}
	return ""
}
