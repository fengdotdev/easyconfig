package disk_test

import (
	"os"
	"testing"

	"github.com/fengdotdev/easyconfig/easyconfig/disk"
	"github.com/stretchr/testify/assert"
)

func TestAssertFileExist(t *testing.T) {
	filepath:= "tempfile.xyz"
	createfile:= os.Create(filepath)
	
	defer os.Remove("tempfile.xyz")
	assert.True(t, disk.AssertFileExists(filepath),true)	
}