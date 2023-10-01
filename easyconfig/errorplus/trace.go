package errorplus

import (
	"fmt"
	"reflect"
	"runtime"
	"time"
)

func E(e error, i interface{})  error {
	
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	s:= runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
	return fmt.Errorf("error: %s trace: %s time: %s", e,s, timestamp)
}

func EString(e string, i interface{})  error {
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	s:= runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
	return fmt.Errorf("error: %s trace: %s  time: %s", e,s, timestamp)
}


func EE(e error, i interface{})  error {
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	s:= runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
	return fmt.Errorf("error: %s trace: %s  time: %s", e,s, timestamp)
}