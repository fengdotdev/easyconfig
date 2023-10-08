package jsonrepo_test

import (
	"encoding/json"
	"testing"

	"github.com/fengdotdev/easyconfig/easyconfig/jsonrepo"

	"github.com/stretchr/testify/assert"
)


func TestJSON_MapToJSON(t *testing.T){

	somemap := make(map[string]interface{})

	somemap["key1"] = "valA"
	somemap["key2"] = "ValB"
	somemap["key3"] = true


	j := jsonrepo.JSONr{}

	jsondata, err := j.MapToJSON(somemap)

	assert.Equal(t, nil, err)
	assert.Equal(t, true, len(jsondata) > 0)

	resultmap := make(map[string]interface{})

	err = json.Unmarshal(jsondata, &resultmap)

	assert.Equal(t, nil, err)
	assert.Equal(t, "valA", resultmap["key1"])

	
}

