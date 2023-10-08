package easyconfig_test

import (
	"github.com/fengdotdev/easyconfig/easyconfig"
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
func TestGetConfig(t *testing.T) {




	config := easyconfig.GetConfig()
	assert.NotNil(t, config)
	assert.NotEmpty(t, config)
	assert.Equal(t, "test", config["test"])
}


*/
//PENDING
//WORKING
func TestIsValidPath(t *testing.T) {
	assert.True(t, easyconfig.IsValidPath("test"))
	assert.True(t, easyconfig.IsValidPath("test/test"))
	assert.True(t, easyconfig.IsValidPath("test/test/test"))
	assert.True(t, easyconfig.IsValidPath("test/test/test/test"))
}


func TestCreateJsonDataFromMap(t *testing.T) {

	map1 := map[string]interface{}{"LLAVE": "VALOR"}
	data, err := easyconfig.CreateJsonDataFromMap(map1)
	assert.Nil(t, err)
	assert.NotNil(t, data)
	assert.NotEmpty(t, data)
	assert.Equal(t, "{\"LLAVE\":\"VALOR\"}", string(data))
}



func TestSimpleIDGenerator(t *testing.T) {
	id1 := easyconfig.SimpleIDGenerator("test")
	assert.NotNil(t, id1)
	assert.NotEmpty(t, id1)
	biggerOrEqual := len(id1) >= 32
	assert.True(t, biggerOrEqual)
	assert.Contains(t, id1, "test")

	id2 := easyconfig.SimpleIDGenerator("test")
	assert.NotNil(t, id2)

	assert.NotEqual(t, id1, id2)

}


