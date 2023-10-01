package jsonrepo

import (
	"easyconfig/pkg/errorplus"
	"encoding/json"
)

//TESTED
//MapToJSON converts a map[string]interface{} to a json in []byte
func (j *JSONr) MapToJSON(m map[string]interface{}) ([]byte, error) {

	//STEP I
	jsondata , err:= json.Marshal(m)

	if err != nil {
		return nil, errorplus.EE(err, j.MapToJSON)
	}

	return jsondata, nil
}