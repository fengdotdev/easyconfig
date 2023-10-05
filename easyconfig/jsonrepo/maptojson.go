package jsonrepo

import (
	"encoding/json"
	"github.com/fengdotdev/goerrorsplus/e"
)

//TESTED
//MapToJSON converts a map[string]interface{} to a json in []byte
func (j *JSONr) MapToJSON(m map[string]interface{}) ([]byte, error) {

	//STEP I
	jsondata , err:= json.Marshal(m)

	if err != nil {
		return nil, e.E(err, "error marshalling json", []string{"jsonrepo", "MapToJSON"}, j.MapToJSON)
	}

	return jsondata, nil
}