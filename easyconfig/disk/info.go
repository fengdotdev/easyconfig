package disk

import (
	"io/fs"
	"os"

	"github.com/fengdotdev/goerrorsplus/e"
)

//TESTME
func GetInfo(path string)(fileinfo fs.FileInfo,err *e.ErrorPlus) {

	defer func() {
		if r := recover(); r != nil {
			re := r.(error)
			err = e.E(re, "error at getting the file info", []string{"disk-package","recover"} , GetInfo, path)
		}
	}()
	fileinfo, er:= os.Stat(path)
	if err != nil {
		return nil, e.E(er, "error at getting the file info", []string{"disk-package"}, GetInfo, path)	
	}

	return fileinfo, nil
}


//TESTME
func GetInfoUnsecure(path string) fs.FileInfo {
	fileinfo, er := os.Stat(path)
	if er != nil {
		return nil
	}
	return fileinfo
}


