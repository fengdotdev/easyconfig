package disk_test

import (
	"os"
	"testing"

	"github.com/fengdotdev/easyconfig/easyconfig/disk"
	"github.com/stretchr/testify/assert"
)

func TestAssertFileExist(t *testing.T) {
	filepath:= "tempfile.xyz"

	assert.False(t, disk.AssertFileExists(filepath),false)
	createfile,err:= os.Create(filepath)
	if err != nil {
		t.Errorf("Error creating file: %v", err)
	}
	createfile.Close()
	assert.True(t, disk.AssertFileExists(filepath),true)
	os.Remove("tempfile.xyz")
	assert.False(t, disk.AssertFileExists(filepath),false)


	// create a directory
	dirpath:= "tempdir"
	assert.False(t, disk.AssertDirectoryExist(dirpath),false)
	err= os.Mkdir(dirpath, 0755)
	if err != nil {
		t.Errorf("Error creating directory: %v", err)
	}
	assert.False(t, disk.AssertFileExists(dirpath),false)
	os.Remove(dirpath)
}



func TestAssertDirectoryExist(t *testing.T) {
	dirpath:= "tempdir"

	assert.False(t, disk.AssertDirectoryExist(dirpath),false)
	err:= os.Mkdir(dirpath, 0755)
	if err != nil {
		t.Errorf("Error creating directory: %v", err)
	}
	assert.True(t, disk.AssertDirectoryExist(dirpath),true)
	os.Remove(dirpath)
	assert.False(t, disk.AssertDirectoryExist(dirpath),false)

	// create a file
	filepath:= "tempfile.xyz"
	info,err:=  os.Stat(filepath)
	assert.False(t, os.Stat(filepath).,false)
	createfile,err:= os.Create(filepath)
	if err != nil {
		t.Errorf("Error creating file: %v", err)
	}
	createfile.Close()
	assert.False(t, disk.AssertDirectoryExist(filepath),false)
	os.Remove("tempfile.xyz")
}
