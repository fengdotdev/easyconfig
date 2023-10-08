package disk

import "os"

//------------------  assertions  ------------------//

//TESTME
func AssertFileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

//TESTME
func AssertDirectoryExist(path string) bool {
	 _, err := os.Stat(path)
	return  err == nil
}
