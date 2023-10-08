package disk_test

import (
	"os"
	"testing"

	"github.com/fengdotdev/easyconfig/easyconfig/disk"
	"github.com/stretchr/testify/assert"
)

func TestCreateDirectory(t *testing.T) {
	path := "tempdir"
	assert.False(t, disk.AssertDirectoryExist(path), false)
	err := disk.CreateDirectory(path)
	if err != nil {
		t.Errorf("Error creating directory: %v", err)
	}
	assert.True(t, disk.AssertDirectoryExist("tempdir"), true)
	os.Remove("tempdir")
	assert.False(t, disk.AssertDirectoryExist("tempdir"), false)

	failedpath := "tempdir1/tempdir2"

	assert.False(t, disk.AssertDirectoryExist(failedpath), false)
	err = disk.CreateDirectory(failedpath)
	if err != nil {
		assert.Error(t, err, "error at creating the directory")
	}
	assert.False(t, disk.AssertDirectoryExist("tempdir1"), false)
	assert.False(t, disk.AssertDirectoryExist(failedpath), false)
	os.RemoveAll("tempdir1")
}


func TestCreateDirectories(t *testing.T) {
	fatherpath := "tempdir1"
	path := "tempdir1/tempdir2"

	assert.False(t, disk.AssertDirectoryExist(fatherpath), false)
	assert.False(t, disk.AssertDirectoryExist(path), false)

	err := disk.CreateDirectories(path)
	if err != nil {
		t.Errorf("Error creating directory: %v", err)
	}
	assert.True(t, disk.AssertDirectoryExist(fatherpath), true)
	assert.True(t, disk.AssertDirectoryExist(path), true)
	os.RemoveAll("tempdir1")
	assert.False(t, disk.AssertDirectoryExist(path), false)
	assert.False(t, disk.AssertDirectoryExist(fatherpath), false)
}