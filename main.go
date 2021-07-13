package main

import (
	"fmt"
	"flag"
	"regexp"
	"errors"
	"github.com/PuerkitoBio/goquery"
)

type Specification struct {
	url string
	version string
}

func formatOutput(spec []Specification) {
	fmt.Println("+-----+---------+----------------------------------------------------------------------------------+")
	fmt.Println("| No. | Version | URL                                                                              |")
	fmt.Println("+-----+---------+----------------------------------------------------------------------------------+")
	for i := 0; i < len(spec); i++ {
		fmt.Printf("| %3d | %7s | %-80s |\n", i + 1, spec[i].version, spec[i].url)
	}
	fmt.Println("+-----+---------+----------------------------------------------------------------------------------+")
}

func stringSearch(targetString string, reString string) (string, error) {
	re := regexp.MustCompile(reString)
	searchResult := re.FindAllStringSubmatch(targetString, -1)
	if searchResult != nil {
		return searchResult[0][0], errors.New("param is empty")
	}
	return "0", nil
}

func getPage(url string) {
	spec := make([]Specification, 0)

	doc, _ := goquery.NewDocument(url)
	doc.Find("a").Each(func(_ int, s *goquery.Selection) {
		href, _ := s.Attr("href")
		text := s.Text()

		str1, err := stringSearch(href, `http.*.zip`)
		if err != nil {
			str2, err := stringSearch(text, `[0-9]{1,2}\.[0-9]{1,2}\.[0-9]{1,2}`)
			if err != nil {
				spec = append(spec, Specification{str1, str2})
			}
		}
	})
	formatOutput(spec)
}

func main() {
	// url := "https://portal.3gpp.org/desktopmodules/Specifications/SpecificationDetails.aspx?specificationId=849"

	url := flag.String("url", "-", "url")
	flag.Parse()

	if *url != "-" {
		getPage(*url)
	} else {
		fmt.Println("Please specify a valid URL.")
	}
}