package jsonrepo

import (
	"easyconfig/pkg/errorplus"
	"encoding/json"
)


//TESTED
//JSONToMap converts a json in []byte to a map[string]interface{}
func (j *JSONr) JSONToMap(b []byte) (map[string]interface{}, error) {

	var m map[string]interface{}


	if len(b) == 0 {
		return nil, errorplus.ES("empty byte slice", j.JSONToMap)
	}

	err := json.Unmarshal(b, &m)

	if err != nil {
		return nil, errorplus.EE(err, j.JSONToMap)
	}

	return m, nil

}