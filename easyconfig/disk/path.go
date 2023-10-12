package disk

import (
	"fmt"
	"strings"
)

//TESTME
func GetParentsPathDir(path string)(parentName string, er error) {
		
	// windows path
	if strings.Contains(path, "\\") {

		return path[:strings.LastIndex(path, "\\")], nil
	}

	// unix path
	if strings.Contains(path, "/") {
		return path[:strings.LastIndex(path, "/")], nil
	}
	return "", fmt.Errorf("path %s is not a valid path", path)
} //  GetParentsPathDir



// WORKING
// TESTME 
func GetFileNameFromPath(path string) (fileName string, er error) {

	// windows path
	if strings.Contains(path, "\\") {

		return path[strings.LastIndex(path, "\\")+1:], nil
	}

	// unix path
	if strings.Contains(path, "/") {
		return path[strings.LastIndex(path, "/")+1:], nil
	}
	return "", fmt.Errorf("path %s is not a valid path", path)
}// GetFileNameFromPath

//TESTME
func GetFileExtension(path string) (extension string, er error) {
	
	filename , err := GetFileNameFromPath(path)
	if err != nil {
		return "", err
	}

	 if strings.Contains(filename, ".") {
		return filename[strings.LastIndex(filename, ".")+1:], nil
	}
	return "", fmt.Errorf("%s does not contain a file extension", filename)
}// GetFileExtension

//TESTME
func GetFileNameWithoutExtension(path string) string {
		panic("implement me")

	return path[:strings.LastIndex(path, ".")]
}// GetFileNameWithoutExtension

//TESTME
func GetSliceOfDirInPath(path string) []string {
		panic("implement me")

	return strings.Split(path, "/")
}// GetSliceOfDirInPath


//TESTME
//NormalizePath replaces all backslashes with forward slashes for a given path in Unix style
func NormalizePath(path string) string {
	panic("implement me")
	return strings.ReplaceAll(path, "\\", "/")
}// NormalizePath

