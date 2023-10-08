package disk

import (
	"os"
	"github.com/fengdotdev/goerrorsplus/e"
)

//------------------  Dir Ops  ------------------//

//TESTME
func CreateDirectories(path string) (err *e.ErrorPlus) {

	defer func() {
		if r := recover(); r != nil {
			re := r.(error)
			err = e.E(re, "error at creating the directory", []string{"disk-package","recover"} , CreateDirectories, path)
		}
	}()

	if !AssertDirectoryExist(path) {
		//create the directory
		er := os.MkdirAll(path, 0755)
		if er != nil {
			return e.E(er, "error at creating the directories", []string{"disk-package"}, CreateDirectories, path)
		}
	}
	return nil
}

//TESTME
func CreateDirectory (path string) (err*e.ErrorPlus) {
	defer func() {
		if r := recover(); r != nil {
			re := r.(error)
			err = e.E(re, "error at creating the directory", []string{"disk-package","recover"} , CreateDirectory, path)
		}
	}()

	if !AssertDirectoryExist(path) {
		//create the directory
		er := os.Mkdir(path, 0755)
		if er != nil {
			return e.E(er, "error at creating the directory", []string{"disk-package"}, CreateDirectories, path)
		}
	}
	return nil
}