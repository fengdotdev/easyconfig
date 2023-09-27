package easyconfig_test

import (
	"easyconfig/easyconfig"
	"fmt"
	"os"
	"testing"

	"github.com/google/uuid"

	"github.com/stretchr/testify/assert"
)



func TestGetConfig(t *testing.T) {
	path := mockFilePath()
	mockJsonFile(path)
	config := easyconfig.NewConfig(path)
	assert.NotNil(t, config)
	assert.Equal(t, path, config.GetConfigPath())

	configData,_ := config.GetConfigData()
	assert.NotNil(t, configData)
	assert.Equal(t, "value1", configData["key1"])

	configData2,_ := config.GetConfigData("key2")
	assert.NotNil(t, configData2)
	assert.Equal(t, "value2", configData2["key2"])
	defer os.Remove(path)
	
}



func mockFilePath()string {
	return fmt.Sprintf("easyconfig%s.json", uuid.New())
}


func TestMockFilePath(t *testing.T) {
	path := mockFilePath()
	assert.Contains(t, path, "easyconfig")
	assert.Contains(t, path, ".json")
}


func mockJsonFile(path string) {
	s:= `{
		"key1": "value1",
		"key2": "value2"
	}`

	
	os.WriteFile(path, []byte(s), 0644)
}

func TestMockJsonFile(t *testing.T){
	path := mockFilePath()
	mockJsonFile(path)
	assert.FileExists(t, path)
	defer os.Remove(path)
}


func TestValidPath(t *testing.T) {
	noexistentPath := fmt.Sprintf("easyconfig%s.json", uuid.New())
	p1 := easyconfig.IsValidPath(noexistentPath)
	assert.False(t, p1)
	//assert.Nil(t, err)


	path := fmt.Sprintf("easyconfig%s.json", uuid.New())

	os.WriteFile(path, []byte("hello"), 0644)

	p2:= easyconfig.IsValidPath(path)
	assert.True(t, p2)
	defer os.Remove(path)
}