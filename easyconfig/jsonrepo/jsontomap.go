package jsonrepo

import (
	"encoding/json"
	"fmt"
	"github.com/fengdotdev/goerrorsplus/e"
)

//TESTED
//JSONToMap converts a json in []byte to a map[string]interface{}
func (j *JSONr) JSONToMap(b []byte) (map[string]interface{}, error) {

	var m map[string]interface{}


	if len(b) == 0 {
		return nil, e.E(fmt.Errorf("err"), "empty json", []string{"jsonrepo", "JSONToMap"}, j.JSONToMap)
	}

	err := json.Unmarshal(b, &m)

	if err != nil {
		return nil, e.E(err, "error unmarshalling json", []string{"jsonrepo", "JSONToMap"}, j.JSONToMap)
	}

	return m, nil

}