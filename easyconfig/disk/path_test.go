package disk_test

import (
	"testing"

	"github.com/fengdotdev/easyconfig/easyconfig/disk"
	"github.com/stretchr/testify/assert"
)

func TestGetParentDir(t *testing.T) {
	assert.Equal(t, "path", disk.GetParentDir("path/test.txt"))
	assert.Equal(t, "path", disk.GetParentDir("path/test"))
	assert.Equal(t, "path", disk.GetParentDir("path/test/"))
	assert.Equal(t, "path", disk.GetParentDir("path/test"))
	assert.Equal(t, "", disk.GetParentDir("test"))
	assert.Equal(t, "", disk.GetParentDir(""))
}



func TestGetFileName(t *testing.T) {
	assert.Equal(t, "test.txt", disk.GetFileName("path/test.txt"))
	assert.Equal(t, "test", disk.GetFileName("path/test"))
	assert.Equal(t, "test", disk.GetFileName("path/test/"))
	assert.Equal(t, "test", disk.GetFileName("path/test"))
	assert.Equal(t, "test", disk.GetFileName("test"))
	assert.Equal(t, "", disk.GetFileName(""))
}