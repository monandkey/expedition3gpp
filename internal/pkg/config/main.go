/*
This package is for configuration management.
*/
package config

import (
	"fmt"
)

// SelectConfigUser is a function that returns an interface.
func SelectConfigUser() ConfigAction {
	return &baseParams{}
}

// SetParams is a function to set the required parameters.
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

// Load is a function for loading the configuration.
func (b *baseParams) Load() params {
	fileName := getFileName()

	if !(fileExist(fileName)) {
		b.SetParams("HOMEDIR", true, 14400, "HOMEDIR")
		if err := b.Write(); err != nil {
			return b.params
		}
	}
	return configLoad(fileName)
}

// Write is a function for writing the configuration.
func (b *baseParams) Write() error {
	var overWriteFlag string
	fileName := getFileName()

	if fileExist(fileName) {
		fmt.Print("overwrite ? y or n: ")
		fmt.Scan(&overWriteFlag)

		if overWriteFlag != "y" {
			fmt.Println("configurations were not changed")
			return nil
		}
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

// GetConfigName is a function to get the configuration name.
func GetConfigName() string {
	return getFileName()
}
