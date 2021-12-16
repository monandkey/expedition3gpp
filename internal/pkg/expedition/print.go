package expedition

import (
	"fmt"
	"regexp"
	"strings"
)

func formatDisplayAll(value []valueBody) {
	maxUrlLen := maxStringLength(value)
	headerPadding := map[string]string{
		"number":  strings.Repeat("-", 5),
		"version": strings.Repeat("-", 9),
		"url":     strings.Repeat("-", maxUrlLen+2),
	}

	fmt.Printf("+%s+%s+%s+\n", headerPadding["number"], headerPadding["version"], headerPadding["url"])
	fmt.Printf("| No. | Version | URL %s |\n", strings.Repeat(" ", maxUrlLen-4))
	fmt.Printf("+%s+%s+%s+\n", headerPadding["number"], headerPadding["version"], headerPadding["url"])
	for i, v := range value {
		urlLen := urlPadding(maxUrlLen, v.Url)
		fmt.Printf("| %3d | %7s | %s%s |\n", i+1, v.Version, v.Url, urlLen)
	}
	fmt.Printf("+%s+%s+%s+\n", headerPadding["number"], headerPadding["version"], headerPadding["url"])
}

func formatDisplayRelease(value []valueBody, releaseNumber string) {
	count := 1
	maxUrlLen := maxStringLength(value)
	headerPadding := map[string]string{
		"number":  strings.Repeat("-", 5),
		"version": strings.Repeat("-", 9),
		"url":     strings.Repeat("-", maxUrlLen+2),
	}

	fmt.Printf("+%s+%s+%s+\n", headerPadding["number"], headerPadding["version"], headerPadding["url"])
	fmt.Printf("| No. | Version | URL %s |\n", strings.Repeat(" ", maxUrlLen-4))
	fmt.Printf("+%s+%s+%s+\n", headerPadding["number"], headerPadding["version"], headerPadding["url"])
	for _, v := range value {
		isBool := regexp.MustCompile(releaseNumber + `.[0-9]{1,2}.[0-9]{1,2}`).MatchString(v.Version)
		if !(isBool) {
			continue
		}
		urlLen := urlPadding(maxUrlLen, v.Url)
		fmt.Printf("| %3d | %7s | %s%s |\n", count, v.Version, v.Url, urlLen)
		count++
	}
	fmt.Printf("+%s+%s+%s+\n", headerPadding["number"], headerPadding["version"], headerPadding["url"])
}

func formatDisplayVersion(value []valueBody, documentVersion string) {
	count := 1
	maxUrlLen := maxStringLength(value)
	headerPadding := map[string]string{
		"number":  strings.Repeat("-", 5),
		"version": strings.Repeat("-", 9),
		"url":     strings.Repeat("-", maxUrlLen+2),
	}

	fmt.Printf("+%s+%s+%s+\n", headerPadding["number"], headerPadding["version"], headerPadding["url"])
	fmt.Printf("| No. | Version | URL %s |\n", strings.Repeat(" ", maxUrlLen-4))
	fmt.Printf("+%s+%s+%s+\n", headerPadding["number"], headerPadding["version"], headerPadding["url"])
	for _, v := range value {
		isBool := regexp.MustCompile(documentVersion).MatchString(v.Version)
		if !(isBool) {
			continue
		}
		urlLen := urlPadding(maxUrlLen, v.Url)
		fmt.Printf("| %3d | %7s | %s%s |\n", count, v.Version, v.Url, urlLen)
		count++
	}
	fmt.Printf("+%s+%s+%s+\n", headerPadding["number"], headerPadding["version"], headerPadding["url"])
}

func maxStringLength(str []valueBody) int {
	max := 0
	for _, v := range str {
		if max < len(v.Url) {
			max = len(v.Url)
		}
	}
	return max
}

func urlPadding(num int, url string) string {
	padding := num - len(url)
	return strings.Repeat(" ", padding)
}

var marks = []string{"|", "/", "-", "\\"}

func mark(i int) string {
	return marks[i%4]
}

func dot(p int) string {
	d := ""
	for i := 0; i <= (p % 3); i++ {
		d += d + "."
	}
	return d
}

func displayLoading(cancel chan struct{}) {
	cnt := 1000000000
	i := 1
	for {
		select {
		case <-cancel:
			return
		default:
			if i%(cnt/100) == 0 {
				p := i / (cnt / 100)
				fmt.Printf("\r[%s] Downloading %-10s", mark(p), dot(p))
			}
			i++
		}
	}
}
