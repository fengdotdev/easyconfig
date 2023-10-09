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


//TESTED
func GetInfoUnsecure(path string) fs.FileInfo {
	fileinfo, er := os.Stat(path)
	if er != nil {
		return nil
	}
	return fileinfo
}



//TESTME
func FileSize(path string) (int64, *e.ErrorPlus) {
	if !AssertFileExists(path) {
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


