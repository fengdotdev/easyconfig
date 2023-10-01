package errorplus_test

import (
	"easyconfig/pkg/errorplus"
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)


func errorFunc() error {
	return errors.New("test")
}

func TestEE(t *testing.T) {
	
	err := errorFunc()

	ee := errorplus.EE(err,errorFunc)
	fmt.Println(ee)
	assert.NotNil(t, ee)
	assert.Contains(t, ee.Error(), "test")
	assert.Contains(t, ee.Error(), "errorplus_test.errorFunc")
	assert.Contains(t, ee.Error(), "trace:")
}

func TestES(t *testing.T) {
	
	es := errorplus.ES("some err",errorFunc)
	fmt.Println(es)
	assert.NotNil(t, es)
	assert.Contains(t, es.Error(), "some err")
	assert.Contains(t, es.Error(), "errorplus_test.errorFunc")
	assert.Contains(t, es.Error(), "trace:")
}