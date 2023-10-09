package disk_test

import (
	"testing"

	"github.com/fengdotdev/easyconfig/easyconfig/disk"
	"github.com/stretchr/testify/assert"
)

func TestNewDisk(t *testing.T) {
	 assert.Equal(t, 1, 1)
}

func TestStorageComplainsWithInterfaceIDisk(t *testing.T) {
	f := func (i disk.IDisk) {
		assert.NotNil(t, i)
	}

	s := disk.Storage{}
	f(&s) // pass a pointer to s instead of s itself


}

