package main

import (
	"easyconfig/easyconfig"
	"fmt"
	"log"
)

//PENDING
func main() {

	fmt.Println(easyconfig.GetRootAPPPath())
	path := "config.json"
	defaultConfigBulder:= easyconfig.NewConfigBuilder(path)

	// Add some default values

	defaultConfigBulder.Add("port", 8080)
	defaultConfigBulder.Add("host", "localhost")

	// Build the config

	defaultConfig, err := defaultConfigBulder.Build()
	if err != nil {
		panic(err)
	}

	log.Println(defaultConfig.GetConfigData())

	// Save the config	
	err = defaultConfigBulder.SaveConfig()
	if err != nil {
		panic(err)
	}

	config := easyconfig.NewConfig(path)

	fmt.Printf("config data from file %s: ", path)
	fmt.Println(config.GetConfigData())
	}