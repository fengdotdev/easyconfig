package main

import (
	"easyconfig/easyconfig"
	"fmt"
	"log"
	"runtime"
)

//PENDING



func Debug(format string, a ...interface{}) {
    _, file, line, _ := runtime.Caller(1)
    info := fmt.Sprintf(format, a...)

    log.Printf("[cgl] debug %s:%d %v", file, line, info)
}


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

	fmt.Println(defaultConfig.GetConfigData())

	// Save the config	
	err = defaultConfigBulder.SaveConfig()
	if err != nil {
		Debug("error: %s", err)
	}

	config := easyconfig.NewConfig(path)

	fmt.Printf("config data from file %s: ", path)
	fmt.Println(config.GetConfigData())





	}