package disk

import (
	"os"
	"github.com/fengdotdev/goerrorsplus/e"
)

//TESTME
func CheckIfFileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

//TESTME
func CheckIfDirectoryExist(path string) bool {
	 _, err := os.Stat(path)
	return  err == nil
}

//TESTME
func CreateDirectories(path string) (err *e.ErrorPlus) {

	defer func() {
		if r := recover(); r != nil {
			re := r.(error)
			err = e.E(re, "error at creating the directory", []string{"disk-package","recover"} , CreateDirectories, path)
		}
	}()

	if !CheckIfDirectoryExist(path) {
		//create the directory
		er := os.MkdirAll(path, 0755)
		if er != nil {
			return e.E(er, "error at creating the directory", []string{"disk-package"}, CreateDirectories, path)
		}
	}

	return nil
}

//TESTME
func FileSize(path string) (int64, *e.ErrorPlus) {
	if !CheckIfFileExists(path) {
		return 0, e.E(nil, "file does not exist", []string{"disk-package"}, FileSize, path)
	}

	file, er := os.Open(path)
	if er != nil {
		return 0, e.E(er, "error at opening the file", []string{"disk-package"}, FileSize, path)
	}

	defer file.Close()

	stat, er := file.Stat()
	if er != nil {
		return 0, e.E(er, "error at getting the file stats", []string{"disk-package"}, FileSize, path)
	}

	return stat.Size(), nil
}

//TESTME
func  FileSizeUnSecure(path string) int64 {
	file, er := os.Open(path)
	if er != nil {
		return 0
	}

	defer file.Close()

	stat, er := file.Stat()
	if er != nil {
		return 0
	}

	return stat.Size()
}