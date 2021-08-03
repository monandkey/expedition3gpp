package expedition3gpp

import (
	"os"
	"fmt"
	"strconv"
	"runtime"
	"io/ioutil"
)

type initConfig struct {
	strageLocation     string
	cacheEnable        bool
	cacheRetentionTime int
	cacheLocation      string
}

type InitConfig struct {
	StrageLocation     string
	CacheEnable        bool
	CacheRetentionTime int
	CacheLocation      string
}

func InitializeConfig() {

	initConfig := initConfig{
		strageLocation: "HOMEDIR",
		cacheEnable: true,
		cacheRetentionTime: 14400,
		cacheLocation: "HOMEDIR",
	}

	homeDir, _ := os.UserHomeDir()
	var fileName *string

	if runtime.GOOS == "windows" {
		str := homeDir + "\\" + ".expedition3gpp.yml"
		fileName = &str
	
	} else if runtime.GOOS == "linux" {
		str := homeDir + "/" + ".expedition3gpp.yml"
		fileName = &str
	
	} else {
		fmt.Println("Your OS is not supported.")
		os.Exit(1)
	}

	fp, err := os.Create(*fileName)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer fp.Close()

	data := [] string{"strageLocation: " + initConfig.strageLocation + "\n",
					  "cacheEnable: " + strconv.FormatBool(initConfig.cacheEnable) + "\n",
					  "cacheRetentionTime: " + strconv.Itoa(initConfig.cacheRetentionTime) + "\n",
					  "cacheLocation: " + initConfig.cacheLocation + "\n"}

	writeConfig(data, *fileName)
}

func writeConfig(data []string, fileName string) {
	b := []byte{}
	for _, line := range data {
		ll := []byte(line)
		for _, l := range ll {
			b = append(b, l)
		}
	}

	err := ioutil.WriteFile(fileName, b, 0666)
	if err != nil {
		fmt.Println(os.Stderr, err)
		os.Exit(1)
	}
}

func ExistInitConfig() bool {
	homeDir, _ := os.UserHomeDir()
	fileName := homeDir + "/" + ".expedition3gpp.yml"
	_, err := os.Stat(fileName)
	return os.IsNotExist(err)
}
