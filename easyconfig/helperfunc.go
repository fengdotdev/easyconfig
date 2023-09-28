package easyconfig

import (
	"encoding/json"
	"os"
)

func GetConfig(path string) (map[string]interface{}, error) {
	file, err := ReadFromFile(path)
	if err != nil {
		return nil, err
	}


	data, err := CreateMapFromJsonData(file)
	if err != nil {
		return nil, err
	}
	return data, nil
}


func IsValidPath(path string) bool {
	_, err := os.Stat(path)
	
	return err== nil
}


func CreateJsonDataFromMap(data map[string]interface{}) ([]byte, error) {
	return json.Marshal(data)
}

func CreateMapFromJsonData(jsondata []byte) (map[string]interface{}, error) {
	var result map[string]interface{}
	err := json.Unmarshal(jsondata, &result)
	return result, err
}


func WriteToFile(path string, data []byte) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.Write(data)
	return err
}

func ReadFromFile(path string) ([]byte, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	
	var data []byte
	_, err = file.Read(data)
	return data, err
}