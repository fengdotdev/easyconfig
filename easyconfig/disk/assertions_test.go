package disk_test

import (
	"os"
	"testing"

	"github.com/fengdotdev/easyconfig/easyconfig/disk"
	"github.com/stretchr/testify/assert"
)

const (


	someAbsoluteUnixPath = "/home/user/somefile.txt" 
	someRelativeUnixPath = "home/user/somefile.txt" 
	someAbsoluteWinPath = "C:\\Users\\user\\somefile.txt" 
	someRelativeWinPath = "Users\\user\\somefile.txt" 
	someWinServerPath = "\\\\server\\share\\somefile.txt" // server path
	someWebResourceUrl = "http://www.example.com/somefile.txt"
	someWebResourceUrl2 = "www.example.com/somefile.txt"
	someWebResourceUrl3 = "./somefile.txt"
	someWebUrl = "http://www.example.com"
	someWebUrl2 = "www.example.com"
	someWebUrl3 = "./"
	someWebUrl4 = "192.168.1.1"
	someWebUrl5 =  "192.168.1.1:8080"
	someFilePath = "somefile.txt" // other path
	someFilePath2 = "somefile" // other path
	someFilePath3 = "/somefile.txt" // other path  ???
	
	someDotPath = "."
	someDotDotPath = ".."
	someEmptyPath = ""
)

func Test_AssertFileExist_default(t *testing.T) {
	filepath:= "tempfile.xyz"

	assert.False(t, disk.AssertFileExists(filepath),false)
	createfile,err:= os.Create(filepath)
	if err != nil {
		t.Errorf("Error creating file: %v", err)
	}
	createfile.Close()
	assert.True(t, disk.AssertFileExists(filepath),true)
	os.Remove("tempfile.xyz")
	assert.False(t, disk.AssertFileExists(filepath),false)


	// create a directory
	dirpath:= "tempdir"
	assert.False(t, disk.AssertDirectoryExist(dirpath),false)
	err= os.Mkdir(dirpath, 0755)
	if err != nil {
		t.Errorf("Error creating directory: %v", err)
	}
	assert.False(t, disk.AssertFileExists(dirpath),false)
	os.Remove(dirpath)
}

func Test_AssertDirectoryExist_default(t *testing.T) {
	dirpath:= "tempdir"

	_, er := os.Stat(dirpath)
	assert.True(t, os.IsNotExist(er), "directory should not exist")
	
	assert.False(t, disk.AssertDirectoryExist(dirpath),false)
	err:= os.Mkdir(dirpath, 0755)
	if err != nil {
		t.Errorf("Error creating directory: %v", err)
	}
	assert.True(t, disk.AssertDirectoryExist(dirpath),true)
	os.Remove(dirpath)
	assert.False(t, disk.AssertDirectoryExist(dirpath),false)

	// create a file
	filepath:= "tempfile.xyz"
	assert.False(t, disk.AssertFileExists(filepath),false)
	createfile,err:= os.Create(filepath)
	if err != nil {
		t.Errorf("Error creating file: %v", err)
	}
	createfile.Close()
	assert.False(t, disk.AssertDirectoryExist(filepath),false)
	os.Remove("tempfile.xyz")

}

//PURE
func Test_AssertIfPathIsAbsolute_default(t *testing.T){

	
	assert.True(t, disk.AssertIfPathIsAbsolute(someAbsoluteUnixPath),true)
	assert.False(t, disk.AssertIfPathIsAbsolute(someRelativeUnixPath),false)
	assert.True(t, disk.AssertIfPathIsAbsolute(someAbsoluteWinPath),true)
	assert.False(t, disk.AssertIfPathIsAbsolute(someRelativeWinPath),false)
	assert.True(t, disk.AssertIfPathIsAbsolute(someWinServerPath),true)
	assert.False(t, disk.AssertIfPathIsAbsolute(someWebResourceUrl),false)
	assert.False(t, disk.AssertIfPathIsAbsolute(someWebResourceUrl2),false)
	assert.False(t, disk.AssertIfPathIsAbsolute(someWebResourceUrl3),false)
	assert.False(t, disk.AssertIfPathIsAbsolute(someWebUrl),false)
	assert.False(t, disk.AssertIfPathIsAbsolute(someWebUrl2),false)
	assert.False(t, disk.AssertIfPathIsAbsolute(someWebUrl3),false)
	assert.False(t, disk.AssertIfPathIsAbsolute(someWebUrl4),false)
	assert.False(t, disk.AssertIfPathIsAbsolute(someWebUrl5),false)
	assert.False(t, disk.AssertIfPathIsAbsolute(someFilePath),false)
	assert.False(t, disk.AssertIfPathIsAbsolute(someFilePath2),false)
	assert.False(t, disk.AssertIfPathIsAbsolute(someFilePath3),false)
	assert.False(t, disk.AssertIfPathIsAbsolute(someDotPath),false)
	assert.False(t, disk.AssertIfPathIsAbsolute(someDotDotPath),false)
	assert.False(t, disk.AssertIfPathIsAbsolute(someEmptyPath),false)
}