package disk_test

import (
	"testing"

	"github.com/fengdotdev/easyconfig/easyconfig/disk"
	"github.com/stretchr/testify/assert"
)

//PURE
//
func Test_GetParentsPathDir_default(t *testing.T) {

	//file only with extension
	fileOnly:= "path_test.go"
	parentDir, err := disk.GetParentsPathDir(fileOnly)
	assert.NotNil(t, err)
	assert.Equal(t, "", parentDir)

	//file only no extension
	fileOnly = "path_test"
	parentDir, err = disk.GetParentsPathDir(fileOnly)
	assert.NotNil(t, err)
	assert.Equal(t, "", parentDir)



	//win relative path to dir
	winRelativePath := ".\\disk\\"
	// windows root

	// windows path with parents
	
	windowsFilePath:= "C:\\Users\\joe\\go\\src\\github.com\\easyconfig\\easyconfig\\disk\\path_test.go"
	parentDir, err := disk.GetParentsPathDir(windowsFilePath)
	if err != nil {
		t.Errorf("Error getting parent dir: %v", err)
	}
	assert.Equal(t, "C:\\Users\\joe\\go\\src\\github.com\\easyconfig\\easyconfig\\disk", parentDir)	



	// unix path with parents
	unixFilePath:= "/home/jow/go/src/github.com/easyconfig/easyconfig/disk/path_test.go"
	parentDir, err = disk.GetParentsPathDir(unixFilePath)
	if err != nil {
		t.Errorf("Error getting parent dir: %v", err)
	}
	assert.Equal(t, "/home/jow/go/src/github.com/easyconfig/easyconfig/disk", parentDir)


}

//PURE
//TESTED confidence minimal
func Test_GetFileNameFromPath_default(t *testing.T) {

	// file with  windowspath and extension
	windowsFilePath:= "C:\\Users\\jow\\go\\src\\github.com\\easyconfig\\easyconfig\\disk\\path_test.go"
	filename, err := disk.GetFileNameFromPath(windowsFilePath)
	if err != nil {
		t.Errorf("Error getting filename: %v", err)
	}
	assert.Equal(t, "path_test.go", filename)	


	// file with unix path and extension
	unixFilePath:= "/home/jow/go/src/github.com/easyconfig/easyconfig/disk/path_test.go"
	filename, err = disk.GetFileNameFromPath(unixFilePath)
	if err != nil {
		t.Errorf("Error getting filename: %v", err)
	}
	assert.Equal(t, "path_test.go", filename)

	// file with windows path and no extension
	windowsFilePath = "C:\\Users\\jow\\go\\src\\github.com\\easyconfig\\easyconfig\\disk\\path_test"
	filename, err = disk.GetFileNameFromPath(windowsFilePath)
	assert.NotNil(t, err)
	assert.Equal(t, "", filename)

	// file with unix path and no extension

	unixFilePath = "/home/jow/go/src/github.com/easyconfig/easyconfig/disk/path_test"
	filename, err = disk.GetFileNameFromPath(unixFilePath)
	assert.NotNil(t, err)
	assert.Equal(t, "", filename)

	// file with extension but no path
	filename, err = disk.GetFileNameFromPath("path_test.go")
	assert.Nil(t, err)
	assert.Equal(t, "path_test.go", filename)

	// no file unix path
	unixFilePath = "/home/jow/go/src/github.com/easyconfig/easyconfig/disk/"
	filename, err = disk.GetFileNameFromPath(unixFilePath)
	assert.NotNil(t, err)
	assert.Equal(t, "", filename)

}//Test_GetFileNameFromPath_default


func Test_GetFileExtension_default(t *testing.T) {

	// file with  windowspath and extension
	windowsFilePath:= "C:\\Users\\jow\\go\\src\\github.com\\easyconfig\\easyconfig\\disk\\path_test.go"
	extension, err := disk.GetFileExtension(windowsFilePath)
	if err != nil {
		t.Errorf("Error getting extension: %v", err)
	}
	assert.Equal(t, "go", extension)	

	// file with unix path and extension
	unixFilePath:= "/home/jow/go/src/github.com/easyconfig/easyconfig/disk/path_test.go"
	extension, err = disk.GetFileExtension(unixFilePath)
	if err != nil {
		t.Errorf("Error getting extension: %v", err)
	}
	assert.Equal(t, "go", extension)

	// file with windows path and no extension
	windowsFilePath = "C:\\Users\\jow\\go\\src\\github.com\\easyconfig\\easyconfig\\disk\\path_test"
	extension, err = disk.GetFileExtension(windowsFilePath)
	assert.NotNil(t, err)
	assert.Equal(t, "", extension)

	// file with unix path and no extension
	unixFilePath = "/home/jow/go/src/github.com/easyconfig/easyconfig/disk/path_test"
	extension, err = disk.GetFileExtension(unixFilePath)
	assert.NotNil(t, err)
	assert.Equal(t, "", extension)

	// no file unix path
	unixFilePath = "/home/jow/go/src/github.com/easyconfig/easyconfig/disk/"
	extension, err = disk.GetFileExtension(unixFilePath)
	assert.NotNil(t, err)
	assert.Equal(t, "", extension)

	// no file windows path
	windowsFilePath = "C:\\Users\\jow\\go\\src\\github.com\\easyconfig\\easyconfig\\disk\\"
	extension, err = disk.GetFileExtension(windowsFilePath)
	assert.NotNil(t, err)
	assert.Equal(t, "", extension)

	// file with extension but no path
	extension, err = disk.GetFileExtension("path_test.go")
	assert.Nil(t, err)
	assert.Equal(t, ".go", extension)
}// Test_GetFileExtension_default