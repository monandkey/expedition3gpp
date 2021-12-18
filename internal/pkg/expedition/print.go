package expedition

import (
	"fmt"
	"regexp"
	"strings"
)

// formatDisplay is a function to display the searched information.
func formatDisplay(value []valueBody, reString string) {
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
		isBool := regexp.MustCompile(reString).MatchString(v.Version)
		if !(isBool) {
			continue
		}
		urlLen := urlPadding(maxUrlLen, v.Url)
		fmt.Printf("| %3d | %7s | %s%s |\n", count, v.Version, v.Url, urlLen)
		count++
	}
	fmt.Printf("+%s+%s+%s+\n", headerPadding["number"], headerPadding["version"], headerPadding["url"])
}

// formatDisplayAll is a function to call formatDisplay.
func formatDisplayAll(value []valueBody) {
	formatDisplay(value, "")
}

// formatDisplayRelease is a function to call formatDisplay.
func formatDisplayRelease(value []valueBody, releaseNumber string) {
	var reString string = `^` + releaseNumber + `\.[0-9]{1,2}\.[0-9]{1,2}`
	formatDisplay(value, reString)
}

// formatDisplayVersion is a function to call formatDisplay.
func formatDisplayVersion(value []valueBody, documentVersion string) {
	formatDisplay(value, documentVersion)
}

// maxStringLength is a function to get the maximum number of characters in a url.
func maxStringLength(str []valueBody) int {
	max := 0
	for _, v := range str {
		if max < len(v.Url) {
			max = len(v.Url)
		}
	}
	return max
}

// urlPadding is a function for padding to the maximum number of characters.
func urlPadding(num int, url string) string {
	padding := num - len(url)
	return strings.Repeat(" ", padding)
}

// marks is a variable that defines the characters to be displayed during search
var marks = []string{"|", "/", "-", "\\"}

// mark is a variable that returns one character from marks
func mark(i int) string {
	return marks[i%4]
}

// dot is a function to display the dot
func dot(p int) string {
	d := ""
	for i := 0; i <= (p % 3); i++ {
		d += d + "."
	}
	return d
}

// displayLoading is a function to display loading on search.
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
