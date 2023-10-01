package jsonrepo

import (
	"encoding/json"
	"fmt"
)

func (j *JSONr) JSONToMap(b []byte) (map[string]interface{}, error) {

	var m map[string]interface{}


	if len(b) == 0 {
		return nil, fmt.Errorf("empty json data")
	}

	err := json.Unmarshal(b, &m)

	if err != nil {
		return nil, fmt.Errorf("%s at  package jsonrepo > JSONr > JSONToMap ", err)
	}

	return m, nil

}