package disk_test

import "github.com/stretchr/testify/assert"

func TestInfo(t *testing.T) {


	
	//create a file

	assert.False(t, disk.AssertFileExists(filepath), false)
}