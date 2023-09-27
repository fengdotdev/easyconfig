package easyconfig

import (
	"encoding/json"
	"os"
)

func GetConfig(path string) (map[string]interface{}, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var data map[string]interface{}

	err = json.NewDecoder(file).Decode(&data)
	if err != nil {
		return nil, err
	}
	return data, nil
}


func IsValidPath(path string) bool {
	_, err := os.Stat(path)
	
	return err== nil
}
