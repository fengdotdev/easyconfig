package disk

import "strings"

//TESTME
func GetParentDir(path string) string {
		panic("implement me")

	return path[:strings.LastIndex(path, "/")]
}

//TESTME
func GetFileName(path string) string {
		panic("implement me")

	return path[strings.LastIndex(path, "/")+1:]
}

//TESTME
func GetFileExtension(path string) string {
		panic("implement me")

	return path[strings.LastIndex(path, ".")+1:]
}

//TESTME
func GetFileNameWithoutExtension(path string) string {
		panic("implement me")

	return path[:strings.LastIndex(path, ".")]
}

//TESTME
func GetSliceOfDirInPath(path string) []string {
		panic("implement me")

	return strings.Split(path, "/")
}


//TESTME
//NormalizePath replaces all backslashes with forward slashes for a given path in Unix style
func NormalizePath(path string) string {
	panic("implement me")
	return strings.ReplaceAll(path, "\\", "/")
}

