package disk_test

import (
	"testing"
		
	"github.com/stretchr/testify/assert"
)

func TestAssertFileExist(t *testing.T) {
	assert.True(t, disk.AssertFileExist("assertions.go"))	
}