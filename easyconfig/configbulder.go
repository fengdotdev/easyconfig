package easyconfig

import (
	"encoding/json"
	"fmt"
	"os"
)
type ConfigBuilder struct {
	path string
	data map[string]interface{}
}

func NewConfigBuilder(path string) *ConfigBuilder {
	return &ConfigBuilder{
		path: path,
		data: make(map[string]interface{}),
	}
}

func (c *ConfigBuilder) Add(key string, value interface{}) {
	c.data[key] = value
}

func (c *ConfigBuilder) Build() (*Config, error) {

	if c.IsConfigValid() {
		return &Config{
			path: c.path,
			data: c.data,
		}, nil
	}

	return nil, fmt.Errorf("invalid config")
}

func (c *ConfigBuilder) SaveConfig(overwrite ...bool) error {

	allowOverwrite := len(overwrite) == 1 && overwrite[0]

	//make json string

	jsondata, err := c.createJsonData()
	if err != nil {
		return err
	}

	// check if file exist
	if !allowOverwrite{
		checkFileExist := IsValidPath(c.path)
		if checkFileExist {
			return fmt.Errorf("file exist")
		}
	}

	// write to file

	file, err := os.Create(c.path)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write	(jsondata)

	return nil
}


func (c *ConfigBuilder) createJsonData() ([]byte, error) {
		jsonData, err := json.Marshal(c.data)
		if err != nil {
			return nil, err
		}
	return jsonData, nil
}

// TODO: implement
func (c *ConfigBuilder) IsConfigValid() bool {
	return true
}