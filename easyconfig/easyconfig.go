package easyconfig

import (
	"encoding/json"
	"fmt"
	"os"
)

func Dummy() {
	fmt.Println("dummy")

}


type Config struct {
	path string
	data map[string]interface{}
}

func NewConfig(pathToJson string) *Config {

	if !IsValidPath(pathToJson) {
		panic("Invalid path")
		}
	

	d,err := GetConfig(pathToJson)
	if err != nil {
		panic(err)
	}
	return &Config{
		path: pathToJson,
		data: d,
	}
}



func (c *Config) GetConfigPath() string {
	return c.path
}

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






