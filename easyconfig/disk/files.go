package disk

import (
	"fmt"
	"os"

	"github.com/fengdotdev/goerrorsplus/e"
)

//------------------  File CRUD  ------------------//

// TESTED
func CreateFile(path string, b []byte) (err *e.ErrorPlus) {
	defer func() {
		if r := recover(); r != nil {
			re := r.(error)
			err = e.E(re, "error at creating the file",[]string{"disk-package","recover"} , CreateFile, path)
		}
	}()

	if AssertFileExists(path) {
		return e.E(nil, "file already exists", []string{"disk-package"}, CreateFile, path)
	}

	file, er := os.Create(path)
	if er != nil {
		return e.E(er, "error at creating the file", []string{"disk-package"}, CreateFile, path)	
	}

	defer file.Close()

	_, er = file.Write(b)
	if err != nil {
		return e.E(er, "error at writing to the file", []string{"disk-package"}, CreateFile, path)
	}

	return nil
}


//TESTME
func  UpdateFile(path string, b []byte) *e.ErrorPlus {
	
	if !AssertFileExists(path) {
		return e.E(nil, "file does not exist", []string{"disk-package"}, UpdateFile, path)
	}

	file, er := os.OpenFile(path, os.O_RDWR, 0644)
	if er != nil {
		return e.E(er, "error at opening the file", []string{"disk-package"}, UpdateFile, path)
	}

	defer file.Close()

	_, er = file.Write(b)

	if er != nil {
		return e.E(er, "error at writing to the file", []string{"disk-package"}, UpdateFile, path)
	}
	
	return nil
}


//TESTME
func DeleteFile(path string) (err *e.ErrorPlus) {

	if !AssertFileExists(path) {
		return e.E(fmt.Errorf("file not exist"), "Error at Deleting the file", []string{"disk-package"}, DeleteFile, path)
	}

	er := os.Remove(path)
	if er != nil {
		return e.E(er, "error at deleting the file", []string{"disk-package"}, DeleteFile, path)
	}
	return nil
}


// TESTME
// read a small file
func ReadFile(path string) (b []byte, err *e.ErrorPlus) {

	defer func() {
		if r := recover(); r != nil {
			re := r.(error)
			err = e.E(re, "error at reading the file", []string{"disk-package","recover"} , ReadFile, path)
		}
	}()

	if !AssertFileExists(path) {
		return nil, e.E(nil, "file does not exist", []string{"disk-package"}, ReadFile, path)
	}


	data, er := os.ReadFile(path)
	if er != nil {
		return nil, e.E(er, "error at reading file", []string{"disk-package"},ReadFile, path)
	}
	return data, nil
}


//TESTME
//LOOKUP
//Read a large file
func  ReadLargeFile(path string,b chan []byte, err chan *e.ErrorPlus){

	defer func() {
		if r := recover(); r != nil {
			re := r.(error)
			b <- nil
			 err <- e.E(re, "error at reading the file", []string{"disk-package","recover"} , ReadLargeFile, path)
		}
	}()

	if !AssertFileExists(path) {
		b <- nil
		err <- e.E(nil, "file does not exist", []string{"disk-package"}, ReadLargeFile, path)
	}

	file, er := os.Open(path)
	if er != nil {
		b <- nil
		err <- e.E(er, "error at opening the file", []string{"disk-package"}, ReadLargeFile, path)
	}

	defer file.Close()

	stat, er := file.Stat()
	if er != nil {
		b <- nil
		err <- e.E(er, "error at getting the file stats", []string{"disk-package"}, ReadLargeFile, path)
	}

	data := make([]byte, stat.Size())
	_, er = file.Read(data)
	if er != nil {
		b <- nil
		err <- e.E(er, "error at reading the file", []string{"disk-package"}, ReadLargeFile, path)
	}
	b <- data
}

// TESTME
// PENDING
func StreamFile(path string)(output chan []byte,err * e.ErrorPlus){
	panic("not implemented")
}


