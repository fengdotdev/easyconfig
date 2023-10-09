package disk_test

import (
	"os"
	"testing"

	"github.com/fengdotdev/easyconfig/easyconfig/disk"
	"github.com/stretchr/testify/assert"
)

//PURE
func TestCreateFile(t *testing.T) {

	//create a file
	filepath := "test.txt"
	err := disk.CreateFile(filepath, []byte("test"))
	assert.Nil(t, err)

	//check if file exists

	_, er := os.Stat(filepath)
	if er != nil {
		assert.Fail(t, "file should exist")
	}
	os.Remove(filepath)
}