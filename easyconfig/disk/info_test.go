package disk_test

import (
	"os"
	"testing"

	"github.com/fengdotdev/easyconfig/easyconfig/disk"
	"github.com/stretchr/testify/assert"
)

//PURE
func TestInfo(t *testing.T) {

	//create a file
	filepath := "test.txt"
	err:= os.WriteFile(filepath, []byte("test"), 0644)	
	assert.Nil(t, err)

	//get the file info
	fileinfo, err := disk.GetInfo(filepath)
	assert.Nil(t, err)
	assert.NotNil(t, fileinfo, "fileinfo should not be nil")
	assert.False(t ,fileinfo.IsDir(), "fileinfo should not be a directory")
	assert.Equal(t, "test.txt", fileinfo.Name(), "fileinfo should have the same name as the file")	
	os.Remove(filepath)
}