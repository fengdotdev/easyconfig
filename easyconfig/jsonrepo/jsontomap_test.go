package jsonrepo_test

import (
	"encoding/json"
	"testing"

	"github.com/fengdotdev/easyconfig/easyconfig/jsonrepo"

	"github.com/stretchr/testify/assert"
)

var(
	databyte = []byte(`{"host":"localhost","port":8080}`)
)




func TestJSON_JSONToMap(t *testing.T){


	

	somemap := make(map[string]interface{})

	somemap["key1"] = "valA"
	somemap["key2"] = "ValB"
	somemap["key3"] = true
	
	jsondata, err := json.Marshal(somemap)

	assert.Equal(t, nil, err)

	j := jsonrepo.JSONr{}

	resultmap, err := j.JSONToMap(jsondata)

	assert.Equal(t, nil, err)

	assert.Equal(t, resultmap["key1"], somemap["key1"])

}



func TestJSON_JSONToMap2 (t *testing.T){

	j := jsonrepo.JSONr{}

	resultmap, err := j.JSONToMap(databyte)

	assert.Equal(t, nil, err)

	assert.Equal(t, resultmap["host"], "localhost")

	assert.Equal(t, resultmap["port"], float64(8080))

}