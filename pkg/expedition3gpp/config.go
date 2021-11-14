package expedition3gpp

import (
	"fmt"
	"io/ioutil"
	"os"
	"runtime"
	"strconv"

	"gopkg.in/yaml.v2"
)

type params struct {
	StrageLocation     string `yaml:"strageLocation"`
	CacheEnable        bool   `yaml:"cacheEnable"`
	CacheRetentionTime int    `yaml:"cacheRetentionTime"`
	CacheLocation      string `yaml:"cacheLocation"`
}

type configPath struct {
	path string
}

func (c configPath) configLoad() params {
	params := params{}
	b, _ := os.ReadFile(c.path)
	yaml.Unmarshal(b, &params)
	return params
}

type disassembledCharacter struct {
	homedir  string
	separate string
	filename string
}

func (d disassembledCharacter) stringJoin() configPath {
	c := configPath{path: d.homedir + d.separate + d.filename}
	return c
}

func getHomedir() string {
	h, err := os.UserHomeDir()
	if err != nil {
		os.Exit(0)
	}
	return h
}

func getSeparate() string {
	switch runtime.GOOS {
	case "windows":
		return "\\"
	case "linux":
		return "/"
	default:
		fmt.Println("Your OS is not support")
		os.Exit(0)
	}
	return ""
}

func getConfigName() string {
	return ".expedition3gpp.yaml"
}

func getConfigParameter() params {
	ds := disassembledCharacter{
		homedir:  getHomedir(),
		separate: getSeparate(),
		filename: getConfigName(),
	}
	return ds.stringJoin().configLoad()
}

func GetConfigParameter() params {
	return getConfigParameter()
}

type InitConfig struct {
	StrageLocation     string
	CacheEnable        bool
	CacheRetentionTime int
	CacheLocation      string
}

type initConfig struct {
	strageLocation     string
	cacheEnable        bool
	cacheRetentionTime int
	cacheLocation      string
}

func InitializeConfig(initConfig *InitConfig) {
	homeDir, _ := os.UserHomeDir()
	var fileName *string

	if runtime.GOOS == "windows" {
		str := homeDir + "\\" + getConfigName()
		fileName = &str

	} else if runtime.GOOS == "linux" {
		str := homeDir + "/" + getConfigName()
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
	data := []string{"strageLocation: " + initConfig.StrageLocation + "\n",
		"cacheEnable: " + strconv.FormatBool(initConfig.CacheEnable) + "\n",
		"cacheRetentionTime: " + strconv.Itoa(initConfig.CacheRetentionTime) + "\n",
		"cacheLocation: " + initConfig.CacheLocation + "\n"}

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
	ds := disassembledCharacter{
		homedir:  getHomedir(),
		separate: getSeparate(),
		filename: getConfigName(),
	}
	f := ds.homedir + ds.separate + ds.filename
	_, err := os.Stat(f)
	return err != nil
}
