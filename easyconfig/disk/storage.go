package disk

import (
	"os"

	"github.com/fengdotdev/goerrorsplus/e"
)

type IDisk interface {
	GetRoot() string

	AssertFileExists(path string) bool
	AssertDirectoryExist(path string) bool

	CreateDirectories(path string) (err *e.ErrorPlus)
	CreateDirectory (path string) (err*e.ErrorPlus)

	CreateFile(path string, b []byte) (err *e.ErrorPlus)
	UpdateFile(path string, b []byte)(err *e.ErrorPlus)
	DeleteFile(path string) (err *e.ErrorPlus)
	ReadFile(path string) (b []byte, err *e.ErrorPlus)
	ReadLargeFile(path string) (b []byte, err *e.ErrorPlus)
}


type Storage struct {
	workingDirectory string
}
//TESTME
func NewStorage(path string) (s *Storage, err *e.ErrorPlus){
	defer func ()  {
		if r := recover(); r != nil {
			re := r.(error)
			s = nil
			err = e.E(re, "error at creating the disk", []string{"disk-package","recover"} , NewStorage, path)
		}
	}()

	if !AssertDirectoryExist(path) {
		//create the directory
		er := os.MkdirAll(path, 0755)
		if er != nil {
			s = nil
			err = e.E(er, "error at creating the directory", []string{"disk-package"}, NewStorage, path)
		}
	}
	return &Storage{workingDirectory: path}, nil
}

//TESTME
func (s *Storage) GetRoot() string {
	return s.workingDirectory
}



