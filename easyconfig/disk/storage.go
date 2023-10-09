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

//TESTME
func (s *Storage) AssertFileExists(path string) bool {
	panic("implement me")
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

//TESTME
func (s *Storage) AssertDirectoryExist(path string) bool {
	panic("implement me")
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false

	}
	return info.IsDir()
}

//TESTME
func (s *Storage) CreateDirectories(path string) (err *e.ErrorPlus) {
	panic("implement me")
	defer func ()  {
		if r := recover(); r != nil {
			re := r.(error)
			err = e.E(re, "error at creating the directories", []string{"disk-package","recover"} , s.CreateDirectories, path)
		}
	}()

	if !AssertDirectoryExist(path) {
		//create the directory
		er := os.MkdirAll(path, 0755)
		if er != nil {
			err = e.E(er, "error at creating the directory", []string{"disk-package"}, s.CreateDirectories, path)
		}
	}
	return nil
}


//TESTME
func (s *Storage) CreateDirectory(path string) (err*e.ErrorPlus) {
	panic("implement me")
	defer func ()  {
		if r := recover(); r != nil {
			re := r.(error)
			err = e.E(re, "error at creating the directory", []string{"disk-package","recover"} , s.CreateDirectory, path)
		}
	}()

	if !AssertDirectoryExist(path) {
		//create the directory
		er := os.Mkdir(path, 0755)
		if er != nil {
			err = e.E(er, "error at creating the directory", []string{"disk-package"}, s.CreateDirectory, path)
		}
	}
	return nil
}

//TESTME
func (s *Storage) CreateFile(path string, b []byte) (err *e.ErrorPlus) {
	panic("implement me")
	defer func ()  {
		if r := recover(); r != nil {
			re := r.(error)
			err = e.E(re, "error at creating the file", []string{"disk-package","recover"} , s.CreateFile, path, b)
		}
	}()

	if AssertFileExists(path) {
		//create the file
		er := os.WriteFile(path, b, 0755)
		if er != nil {
			err = e.E(er, "error at creating the file", []string{"disk-package"}, s.CreateFile, path, b)
		}
	}
	return nil
}

//TESTME
func (s *Storage) UpdateFile(path string, b []byte)(err *e.ErrorPlus) {
	panic("implement me")
	defer func ()  {
		if r := recover(); r != nil {
			re := r.(error)
			err = e.E(re, "error at updating the file", []string{"disk-package","recover"} , s.UpdateFile, path, b)
		}
	}()

	if AssertFileExists(path) {
		//create the file
		er := os.WriteFile(path, b, 0755)
		if er != nil {
			err = e.E(er, "error at updating the file", []string{"disk-package"}, s.UpdateFile, path, b)
		}
	}
	return nil
}

//TESTME
func (s *Storage) DeleteFile(path string) (err *e.ErrorPlus) {
	panic("implement me")
	defer func ()  {
		if r := recover(); r != nil {
			re := r.(error)
			err = e.E(re, "error at deleting the file", []string{"disk-package","recover"} , s.DeleteFile, path)
		}
	}()

	if AssertFileExists(path) {
		//create the file
		er := os.Remove(path)
		if er != nil {
			err = e.E(er, "error at deleting the file", []string{"disk-package"}, s.DeleteFile, path)
		}
	}
	return nil
}

//TESTME
func (s *Storage)ReadFile(path string) (b []byte, err *e.ErrorPlus){
panic("implement me")
}


//TESTME
func (s *Storage)ReadLargeFile(path string) (b []byte, err *e.ErrorPlus){
	panic("implement me")
}