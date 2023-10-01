package jsonrepo

import (
	"easyconfig/easyconfig/errorplus"
	"encoding/json"
)

func (j *JSONr) MapToJSON(m map[string]interface{}) ([]byte, error) {

	//STEP I
	jsondata , err:= json.Marshal(m)

	if err != nil {
		return nil, errorplus.EE(err, j.MapToJSON)
	}

	return jsondata, nil
}