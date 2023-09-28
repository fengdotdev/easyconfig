package main

import (
	"easyconfig/easyconfig"
	"log"
)

func main() {

	path := "config.json"
	defaultConfigBulder:= easyconfig.NewConfigBuilder(path)

	// Add some default values

	defaultConfigBulder.Add("port", 8080)
	defaultConfigBulder.Add("host", "localhost")

	// Build the config

	defaultConfig, err := defaultConfigBulder.Build()
	if err != nil {
		log.Fatal(err)
	}

	log.Println(defaultConfig.GetConfigData())

	// Save the config	
	err = defaultConfigBulder.SaveConfig()
	if err != nil {
		panic(err)
	}

	config := easyconfig.NewConfig(path)

	log.Println(config.GetConfigData())
	}