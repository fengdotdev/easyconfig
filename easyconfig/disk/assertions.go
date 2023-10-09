package disk

import (
	"os"
	"strings"
)

//------------------  assertions  ------------------//

//TESTED
func AssertFileExists(path string) bool {
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

//TESTED
func AssertDirectoryExist(path string) bool {
	 info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
		
	}
	return info.IsDir()
}



//TESTME
func AssertIfPathIsAbsolute(path string) bool {
		panic("implement me")

	return strings.HasPrefix(path, "/")
}

//TESTME
func AssertIfPathIsRelative(path string) bool {
		panic("implement me")

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