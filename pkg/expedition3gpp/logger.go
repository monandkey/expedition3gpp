package expedition3gpp

import (
	"fmt"
	"runtime"
	"github.com/fatih/color"
)

// --------------------------------------------------
// Output
// --------------------------------------------------
func formatOutput(spec []specDocInfo) {
	if runtime.GOOS == "windows" {
		fmt.Println("+-----+---------+----------------------------------------------------------------------------------+")
		fmt.Println("| No. | Version | URL                                                                              |")
		fmt.Println("+-----+---------+----------------------------------------------------------------------------------+")
		for i := 0; i < len(spec); i++ {
			fmt.Printf("| %3d | %7s | %-80s |\n", i + 1, spec[i].version, spec[i].url)
		}
		fmt.Println("+-----+---------+----------------------------------------------------------------------------------+")
		return
	}

	if runtime.GOOS == "linux" {
		fmt.Println("+-----+---------+----------------------------------------------------------------------------------+")
		fmt.Println("| No. | Version | URL                                                                              |")
		fmt.Println("+-----+---------+----------------------------------------------------------------------------------+")
		for i := 0; i < len(spec); i++ {
			fmt.Printf("| %3d | %7s | %-89s |\n", i + 1, spec[i].version, color.HiCyanString(spec[i].url))
		}
		fmt.Println("+-----+---------+----------------------------------------------------------------------------------+")
		return
	}
}

func formatOutputOneVersion(spec []specDocInfo, version string) {
	if runtime.GOOS == "windows" {
		fmt.Println("+-----+---------+----------------------------------------------------------------------------------+")
		fmt.Println("| No. | Version | URL                                                                              |")
		fmt.Println("+-----+---------+----------------------------------------------------------------------------------+")
		for i := 0; i < len(spec); i++ {
			if spec[i].version == version {
				fmt.Printf("| %3d | %7s | %-80s |\n", 1, spec[i].version, spec[i].url)
			}
		}
		fmt.Println("+-----+---------+----------------------------------------------------------------------------------+")
		return
	}

	if runtime.GOOS == "linux" {
		fmt.Println("+-----+---------+----------------------------------------------------------------------------------+")
		fmt.Println("| No. | Version | URL                                                                              |")
		fmt.Println("+-----+---------+----------------------------------------------------------------------------------+")
		for i := 0; i < len(spec); i++ {
			if spec[i].version == version {
				fmt.Printf("| %3d | %7s | %-89s |\n", 1, spec[i].version, color.HiCyanString(spec[i].url))
			}
		}
		fmt.Println("+-----+---------+----------------------------------------------------------------------------------+")
		return
	}
}

func formatOutputYaml(cy cacheYaml) {
	if runtime.GOOS == "windows" {
		fmt.Println("+-----+---------+----------------------------------------------------------------------------------+")
		fmt.Println("| No. | Version | URL                                                                              |")
		fmt.Println("+-----+---------+----------------------------------------------------------------------------------+")
		for i := 0; i < len(cy.Value); i++ {
			fmt.Printf("| %3d | %7s | %-80s |\n", i + 1, cy.Value[i].Version, cy.Value[i].Url)
		}
		fmt.Println("+-----+---------+----------------------------------------------------------------------------------+")
		return
	}

	if runtime.GOOS == "linux" {
		fmt.Println("+-----+---------+----------------------------------------------------------------------------------+")
		fmt.Println("| No. | Version | URL                                                                              |")
		fmt.Println("+-----+---------+----------------------------------------------------------------------------------+")
		for i := 0; i < len(cy.Value); i++ {
			fmt.Printf("| %3d | %7s | %-89s |\n", i + 1, cy.Value[i].Version, color.HiCyanString(cy.Value[i].Url))
		}
		fmt.Println("+-----+---------+----------------------------------------------------------------------------------+")
		return
	}
}

func formatOutputYamlOneVersion(cy cacheYaml, version string) {
	if runtime.GOOS == "windows" {
		fmt.Println("+-----+---------+----------------------------------------------------------------------------------+")
		fmt.Println("| No. | Version | URL                                                                              |")
		fmt.Println("+-----+---------+----------------------------------------------------------------------------------+")
		for i := 0; i < len(cy.Value); i++ {
			if cy.Value[i].Version == version {
				fmt.Printf("| %3d | %7s | %-80s |\n", 1, cy.Value[i].Version, cy.Value[i].Url)
			}
		}
		fmt.Println("+-----+---------+----------------------------------------------------------------------------------+")
		return
	}

	if runtime.GOOS == "linux" {
		fmt.Println("+-----+---------+----------------------------------------------------------------------------------+")
		fmt.Println("| No. | Version | URL                                                                              |")
		fmt.Println("+-----+---------+----------------------------------------------------------------------------------+")
		for i := 0; i < len(cy.Value); i++ {
			if cy.Value[i].Version == version {
				fmt.Printf("| %3d | %7s | %-89s |\n", 1, cy.Value[i].Version, color.HiCyanString(cy.Value[i].Url))
			}
		}
		fmt.Println("+-----+---------+----------------------------------------------------------------------------------+")
		return
	}
}

// --------------------------------------------------
// Display Downloading
// --------------------------------------------------
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
		case <- cancel:
			return

		default:
			if i%(cnt/100) == 0 {
				p := i / (cnt / 100)
				fmt.Printf("\r[ %s ] Downloading %-10s", mark(p), dot(p))
			}
			i++
		}
	}
}
