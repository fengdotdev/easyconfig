package disk

import (
	"fmt"
	"os"

	"github.com/fengdotdev/goerrorsplus/e"
)

type Disk struct {
	workingDirectory string
}
//TESTME
func NewDisk(path string) (d *Disk, err *e.ErrorPlus){
	defer func ()  {
		if r := recover(); r != nil {
			re := r.(error)
			d = nil
			err = e.E(re, "error at creating the disk", []string{"disk-package","recover"} , NewDisk, path)
		}
	}()

	if !CheckIfDirectoryExist(path) {
		//create the directory
		er := os.MkdirAll(path, 0755)
		if er != nil {
			d = nil
			err = e.E(er, "error at creating the directory", []string{"disk-package"}, NewDisk, path)
		}
	}
	return &Disk{workingDirectory: path}, nil
}

//TESTME





func (d *Disk) GetRoot() string {
	return d.workingDirectory
}




// TESTME
func (d *Disk) CreateFile(path string, b []byte) (err *e.ErrorPlus) {
	defer func() {
		if r := recover(); r != nil {
			re := r.(error)
			err = e.E(re, "error at creating the file",[]string{"disk-package","recover"} , d.CreateFile, path)
		}
	}()

	if CheckIfFileExists(path) {
		return e.E(nil, "file already exists", []string{"disk-package"}, d.CreateFile, path)
	}

	file, er := os.Create(path)
	if er != nil {
		return e.E(er, "error at creating the file", []string{"disk-package"}, d.CreateFile, path)	
	}

	defer file.Close()

	_, er = file.Write(b)
	if err != nil {
		return e.E(er, "error at writing to the file", []string{"disk-package"}, d.CreateFile, path)
	}

	return nil
}

func (d *Disk) UpdateFile(path string, b []byte) *e.ErrorPlus {
	
	if !CheckIfFileExists(path) {
		return e.E(nil, "file does not exist", []string{"disk-package"}, d.UpdateFile, path)
	}

	file, er := os.OpenFile(path, os.O_RDWR, 0644)
	if er != nil {
		return e.E(er, "error at opening the file", []string{"disk-package"}, d.UpdateFile, path)
	}

	defer file.Close()

	_, er = file.Write(b)

	if er != nil {
		return e.E(er, "error at writing to the file", []string{"disk-package"}, d.UpdateFile, path)
	}
	
	return nil
}


//TESTME
func (d *Disk) DeleteFile(path string) (err *e.ErrorPlus) {

	if !CheckIfFileExists(path) {
		return e.E(fmt.Errorf("file not exist"), "Error at Deleting the file", []string{"disk-package"}, d.DeleteFile, path)
	}

	er := os.Remove(path)
	if er != nil {
		return e.E(er, "error at deleting the file", []string{"disk-package"}, d.DeleteFile, path)
	}
	return nil
}


// TESTME
// read a small file
func (d *Disk) ReadFile(path string) (b []byte, err *e.ErrorPlus) {

	defer func() {
		if r := recover(); r != nil {
			re := r.(error)
			err = e.E(re, "error at reading the file", []string{"disk-package","recover"} , d.ReadFile, path)
		}
	}()

	if !CheckIfFileExists(path) {
		return nil, e.E(nil, "file does not exist", []string{"disk-package"}, d.ReadFile, path)
	}


	data, er := os.ReadFile(path)
	if er != nil {
		return nil, e.E(er, "error at reading file", []string{"disk-package"}, d.ReadFile, path)
	}
	return data, nil
}


//TESTME
//LOOKUP
//Read a large file
func (d *Disk) ReadLargeFile(path string,b chan []byte, err chan *e.ErrorPlus){

	defer func() {
		if r := recover(); r != nil {
			re := r.(error)
			b <- nil
			 err <- e.E(re, "error at reading the file", []string{"disk-package","recover"} , d.ReadLargeFile, path)
		}
	}()

	if !CheckIfFileExists(path) {
		b <- nil
		err <- e.E(nil, "file does not exist", []string{"disk-package"}, d.ReadLargeFile, path)
	}

	file, er := os.Open(path)
	if er != nil {
		b <- nil
		err <- e.E(er, "error at opening the file", []string{"disk-package"}, d.ReadLargeFile, path)
	}

	defer file.Close()

	stat, er := file.Stat()
	if er != nil {
		b <- nil
		err <- e.E(er, "error at getting the file stats", []string{"disk-package"}, d.ReadLargeFile, path)
	}

	data := make([]byte, stat.Size())
	_, er = file.Read(data)
	if er != nil {
		b <- nil
		err <- e.E(er, "error at reading the file", []string{"disk-package"}, d.ReadLargeFile, path)
	}

	b <- data

}

// TESTME
// PENDING
func (d *Disk) StreamFile(path string)(output chan []byte,err * e.ErrorPlus){
	panic("not implemented")
}





// TESTME
// LOOKUP
func (d *Disk) CreateDirectory(path string) *e.ErrorPlus {
	if d.DirectoryExist(path) {
		return e.E(nil, "directory already exists", []string{"disk-package"}, d.CreateDirectory, path)
	}

	er := os.Mkdir(path, 0755)
	if er != nil {
		return e.E(er, "error at creating the directory", []string{"disk-package"}, d.CreateDirectory, path)
	}
	return nil
}


// TESTME
func (d *Disk) DirectoryExist(path string) bool {
	 _, err := os.Stat(path) 
	return  err == nil
}