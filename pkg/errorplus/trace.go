package errorplus

import (
	"fmt"
	"reflect"
	"runtime"
	"time"
)

//TESTME
// feeded with the fn to return the trace and timestamp
func TimeStampAndRuntimeFN(fn interface{}) string {
	t := time.Now()
	timestamp := t.Format("2006-01-02 15:04:05")
	s:= runtime.FuncForPC(reflect.ValueOf(fn).Pointer()).Name()
	return fmt.Sprintf("trace: %s time: %s", s, timestamp)
}
//TESTME
// ERROR: feeded with the  error and fn to return a error with the trace and timestamp
func EE(e error, fn interface{})  error {
	s := TimeStampAndRuntimeFN(fn)
	return fmt.Errorf("error: %s  %s", e,s,)
}

//TESTME
// ERROR STRING:feeded with the  string error and fn to return a error with the trace and timestamp
func ES(e string, fn interface{})  error {
	s := TimeStampAndRuntimeFN(fn)
	return fmt.Errorf("error: %s  %s ", e,s)
}



type ErrorPlus struct {
	Err error
	Message string
	Tags []string
	FN string
	Args []interface{}
	Trace string
	Time time.Time  
}

// Short hand for NewErrorPlus
func NEP(err error, message string,tags []string, fn interface{}, args ...interface{}) *ErrorPlus {
	errPlus := NewErrorPlus(err, message, tags, fn, args...)
	return errPlus
}


//TESTME
func NewErrorPlus(err error, message string, tags []string,fn interface{}, args ...interface{}) *ErrorPlus {
	e := &ErrorPlus{
		Err: err,
		Tags: tags,
		Message: message,
		FN: runtime.FuncForPC(reflect.ValueOf(fn).Pointer()).Name(),
		Args: args,
		Trace: TimeStampAndRuntimeFN(fn),
		Time: time.Now(),
	}
	return e
}


func (e *ErrorPlus) Error() string {
	err := e.Err.Error()
	return err
}

//short hand for VerboseError
func (e *ErrorPlus)V () error{
	return e.VerboseError()
}


func (e *ErrorPlus) VerboseError() error {

	timestamp := e.Time.Format("2006-01-02 15:04:05")
	err := e.Err.Error()
	message := e.Message
	trace := e.Trace
	fn := e.FN
	args := e.Args

	//format 2006-01-02 15:04:05 error: original-error trace: /something/something/fn.errorFunc args: [1 2 3]
	return fmt.Errorf("%s error: %s  %s fn: %s trace: %s  %s", timestamp, err ,message, fn, trace, args)
}

