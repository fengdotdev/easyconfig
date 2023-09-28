package easyconfig

import (
	"fmt"
)


type Config struct {
	path string
	data map[string]interface{}
}


//TESTME
func NewConfig(pathToJson string, defaultConfig ...Config) *Config {

	existDefaultConfig := len(defaultConfig) == 1

	if len(defaultConfig) > 1 {
		panic("only one default config is allowed")
	}

	if !IsValidPath(pathToJson) {
		panic("Invalid path")
	}
	outPutData := make(map[string]interface{})

	dataFromFile,err := GetConfig(pathToJson)
	if err != nil {
		panic(err)
	}


	if existDefaultConfig {
		dataFromDefaultConfig := defaultConfig[0].data
		for key, value := range dataFromDefaultConfig {
			outPutData[key] = value
		}
	}

	for key, value := range dataFromFile {
			outPutData[key] = value
	}
	return &Config{
		path: pathToJson,
		data: outPutData,
	}
}

//TESTME
func (c *Config) GetConfigPath() string {
	return c.path
}

//TESTME
func (c *Config) GetConfigData(key ...string) (map[string]interface{}, error ){

	if len(key) == 0 {
		return c.data, nil
	}

	if len(key) > 1 {
		return nil, fmt.Errorf("only one key is allowed")
	}

	if val, ok := c.data[key[0]]; ok {
		return map[string]interface{}{key[0]: val}, nil
	}

	return nil, fmt.Errorf("key not found")
}

//TESTME
func (c *Config) SaveConfig(overwrite ...bool)error{

	allowOverwrite := len(overwrite) == 1 && overwrite[0]

	if !allowOverwrite{
		checkFileExist := IsValidPath(c.path)
		if checkFileExist {
			return fmt.Errorf("file exist")
		}
	}

	//make json string
	jsondata, err :=   CreateJsonDataFromMap(c.data)
	//TODO: createJsonData()
	// check if file exist

	// write to file
	err = WriteToFile(c.path, jsondata)
	if err != nil {
		panic(err)
	}
	return nil
}






