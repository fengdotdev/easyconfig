package disk

import (
	"fmt"
	"os"

	"github.com/fengdotdev/goerrorsplus/e"
)

//TESTME
func (d *Disk) Write(b []byte, path string, fileName string,  overwrite ...bool) (outErr *e.ErrorPlus) {

	defer func(){
		if r := recover(); r != nil {
			outErr = e.E(fmt.Errorf("panic"), "panic", []string{"disk", "write"}, d.Write)
		}
	}()

	fullPath := d.FullPath(path, fileName)
	allowOverwrite := false
	pathExists := d.DirectoryExist(path)


	if len(overwrite) == 1 && overwrite[0] {
		allowOverwrite = true
	}


	if d.IsFileExists(fullPath) && !allowOverwrite{
		return e.E(fmt.Errorf("err"), "file already exist", []string{"disk", "write"}, d.Write)
	}
	


	if pathExists == false {
		err := os.MkdirAll(path, 0755)
		if err != nil {
			return e.E(err, "cannot create directory", []string{"disk", "write"}, d.Write)
		}
	}


	file, err := os.Create(fullPath)
	

	if err != nil {
		return e.E(err, "cannot create file", []string{"disk", "write"}, d.Write)
	}

	_, err = file.Write(b)
	if err != nil {
		return e.E(err, "cannot write to file", []string{"disk", "write"}, d.Write)
	}

	defer file.Close()
	return nil
}
