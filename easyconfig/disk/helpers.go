package disk

import (
	"os"
	"strings"
)

//TESTME
func (d *Disk) IsFileExists(path string) bool {
_, err := os.Stat(path)
	return err ==nil
}

//TESTME
func (d *Disk) DirectoryExist(path string)bool {

	_,err := os.Stat(path)
	return err == nil
}

//TESTME
func (d *Disk) FullPath(RawPath,RawFilename string) string {

	fileName :=""

	if strings.HasPrefix(RawFilename, "/") {
		fileName = strings.TrimPrefix(RawFilename, "/")
	} else {
		fileName = RawFilename  
	}

	if strings.HasSuffix(RawPath, "/") {
		return RawPath + fileName
	}

	return RawPath + "/" + fileName
}

//TESTME
func (d *Disk) RootDir() string {
	panic("not implemented")
}
//TESTME
func (d *Disk) DataDir() string {
	panic("not implemented")
}
//TESTME
func (d *Disk) ConfigDir() string {
	panic("not implemented")
}
//TESTME
func (d *Disk) TestDirWR() error {
	panic("not implemented")
}