package disk

import (
	"os"
	"strings"
)

//------------------  assertions  ------------------//

//TESTED
func AssertFileExists(path string) bool {

	exists := false

	//criteria 1: file exists
	info, err := os.Stat(path)
	if err != nil {
		return exists
	}
	exists = true

	//criteria 2: file is not a directory
	if  info.IsDir() {
		exists = false
		return exists
	}

	return exists 
}

//TESTED
func AssertDirectoryExist(path string) bool {

	exists := false

	//criteria 1: file exists
	 info, err := os.Stat(path)
	if err != nil {
		return exists
	}
	exists = true

	//criteria 2: file is a directory
	if  !info.IsDir() {
		exists = false
		return exists
	}
	return exists
}



//TESTED
func AssertIfPathIsAbsolute(path string) bool {

	var isAbsolute bool
	
	//criteria 1: starts with a slash
	if strings.HasPrefix(path, "/") {
		isAbsolute = true
	}

	//criteria 2: starts with a drive letter
	if len(path) >= 2 {
		if path[1] == ':' {
			isAbsolute = true
		}
	}

	//criteria 3: starts with a double backslash
	if len(path) >= 2 {
		if path[0] == '\\' && path[1] == '\\' {
			isAbsolute = true
		}
	}

	return isAbsolute
}

//TESTME
func AssertIfPathIsRelative(path string) bool {
	
	isRelative := false

	//criteria 1: starts with a slash
	if strings.HasPrefix(path, "/") {
		
		
	}

	return !strings.HasPrefix(path, "/")
}

//TESTME
func AssertIfPathIsRoot(path string) bool {
		panic("implement me")

	return path == "/"
}

//TESTME
func AssertIfPathIsDirectory(path string) bool {
		panic("implement me")

	return strings.HasSuffix(path, "/")
}

//TESTME
func AssertIfPathIsFile(path string) bool {
		panic("implement me")

	return !strings.HasSuffix(path, "/")
}

//TESTME
//LOOKUP
// linux hidden files start with a dot
func AssertIfPathIsHidden(path string) bool {
		panic("implement me")

	return strings.HasPrefix(path, ".")
}

//TESTME
func AssertPathIsUnixStyle(path string) bool {
		panic("implement me")

	return strings.Contains(path, "/")
}

//TESTME
func AssertPathIsWindowsStyle(path string) bool {
		panic("implement me")

	return strings.Contains(path, "\\")
}

//TESTME
func AssertPathIsAbsoluteUnixStyle(path string) bool {
		panic("implement me")

	return AssertIfPathIsAbsolute(path) && AssertPathIsUnixStyle(path)
}

//TESTME
func AssertPathIsAbsoluteWindowsStyle(path string) bool {
		panic("implement me")

	return AssertIfPathIsAbsolute(path) && AssertPathIsWindowsStyle(path)
}