package easyconfig

import (
	"easyconfig/easyconfig/jsonrepo"
	"fmt"
	"os"
)
type ConfigBuilder struct {
	path string
	data map[string]interface{}
}

//TESTME NewConfigBuilder
func NewConfigBuilder(path string) *ConfigBuilder {
	return &ConfigBuilder{
		path: path,
		data: make(map[string]interface{}),
	}
}
//TESTME(c *ConfigBuilder) Add
func (c *ConfigBuilder) Add(key string, value interface{}) {
	c.data[key] = value
}

//TESTME (c *ConfigBuilder) Build()
func (c *ConfigBuilder) Build() (*Config, error) {

	// TODO : check if config is valid
	if c.IsConfigValid() {
		return &Config{
			path: c.path,
			data: c.data,
		}, nil
	}

	return nil, fmt.Errorf("invalid config")
}

//TESTME (c *ConfigBuilder) SaveConfig
func (c *ConfigBuilder) SaveConfig(overwrite ...bool) error {

	allowOverwrite := len(overwrite) == 1 && overwrite[0]

	//make json string

	//WORKING
	jsondata, err := c.createJsonData()
	if err != nil {
		return err
	}

	// check if file exist
	if !allowOverwrite{
		
		if CheckIfFileExist(c.path) {
			return fmt.Errorf("file %s exist", c.path)
		}
	}

	// write to file

	file, err := os.Create(c.path)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write	(jsondata)

	return err
}


//TESTME (c *ConfigBuilder) createJsonData
func (c *ConfigBuilder) createJsonData() ([]byte, error) {
		
	json := jsonrepo.JSONr{}

	jsondata, err := json.MapToJSON(c.data)
	if err != nil {
		return nil, err
	}

	return jsondata, nil
}

// TODO: implement (c *ConfigBuilder) IsConfigValid
func (c *ConfigBuilder) IsConfigValid() bool {
	return true
}