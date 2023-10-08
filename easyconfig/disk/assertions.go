package disk

import "os"

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


