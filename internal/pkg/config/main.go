package config

import (
	"errors"
	"fmt"
)

func SelectConfigUser() ConfigAction {
	return &baseParams{}
}

func (b *baseParams) SetParams(
	strageLocation string,
	cacheEnable bool,
	cacheRetentionTime int,
	cacheLocation string,
) {
	b.params.StrageLocation = strageLocation
	b.params.CacheEnable = cacheEnable
	b.params.CacheRetentionTime = cacheRetentionTime
	b.params.CacheLocation = cacheLocation
}

func (b *baseParams) Load() params {
	fileName := getFileName()

	if !(fileExist(fileName)) {
		return b.params
	}
	return configLoad(fileName)
}

func (b *baseParams) Write() error {
	var overWriteFlag string
	fileName := getFileName()

	if fileExist(fileName) {
		fmt.Print("overwrite ? y or n: ")
		fmt.Scan(&overWriteFlag)
	}

	if overWriteFlag != "y" {
		return errors.New("configurations were not changed")
	}

	if err := configWrite(fileName, b.params); err != nil {
		return err
	}

	if overWriteFlag == "y" {
		fmt.Println("Overwrite!!")
	} else {
		fmt.Println("Create config file")
	}
	return nil
}
