package easyconfig

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"math/big"
	"os"
	"time"

	"github.com/google/uuid"
)




func GetRootAPPPath() string {
	dir, _ := os.Getwd()
	return dir
}


//TESTME  GetConfig
func GetConfig(path string) (map[string]interface{}, error) {
	file, err := ReadFromFile(path)
	if err != nil {
		return nil, err
	}


	data, err := CreateMapFromJsonData(file)
	if err != nil {
		return nil, err
	}
	return data, nil
}

//TESTME IsValidPath
func IsValidPath(path string) bool {

	dir, _ := os.Getwd()
	testname:= uuid.New().String()
	testpath := dir + "/" + testname

	// check if the directory exist and can write to the directory
	err := WriteToFile(testpath, []byte("test"))
	if err != nil {
		return false
	}

	if !CheckIfFileExist(testpath) {
		return false
	}

	err = DeleteFile(testpath)
	return err == nil
}

//Ok
func CreateJsonDataFromMap(data map[string]interface{}) ([]byte, error) {
	return json.Marshal(data)
}

//TESTME CreateMapFromJsonData
func CreateMapFromJsonData(jsondata []byte) (map[string]interface{}, error) {
	var result map[string]interface{}
	err := json.Unmarshal(jsondata, &result)
	return result, err
}


//TESTME CheckIfFileExist 
func CheckIfFileExist (path string) bool {
	_, err := os.Stat(path)
	return err== nil
}

//TESTME WriteToFile
func WriteToFile(path string, data []byte) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.Write(data)
	return err
}

//TESTME ReadFromFile
func ReadFromFile(path string) ([]byte, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	
	var data []byte
	_, err = file.Read(data)
	return data, err
}

//TESTME DeleteFile
func DeleteFile(path string) error {
	err := os.Remove(path)
	return err
}

//OK
//this id generator is not safe for Collisions and is only intended for CRUD dummy data
func SimpleIDGenerator(prefix ...string) string {
	
	isPrefix := len(prefix) == 1 && prefix[0] != ""

	timeNow:= time.Now().UnixNano()

	
	//	generate random number between 0 and 9999
	randNumberLimit9999, err := rand.Int(rand.Reader, big.NewInt(9999))
	if err != nil {
		panic(err)
	}

	//	generate random number between 0 and 999
		randNumberLimit999, err := rand.Int(rand.Reader, big.NewInt(999))
	if err != nil {
		panic(err)
	}

	numberProductOf2RandNumbers := big.NewInt(0).Mul(randNumberLimit9999, randNumberLimit999)

	timeNowString := fmt.Sprintf("%x", timeNow)
	numberProductOf2RandNumbersString:= fmt.Sprintf("%x", numberProductOf2RandNumbers)
	hashsha256byte := sha256.Sum256([]byte(timeNowString + numberProductOf2RandNumbersString))
	hashString := fmt.Sprintf("%x", hashsha256byte)

	if isPrefix {
		return prefix[0] + "-" + hashString 
	}


	if checkIfIDExist(hashString){
		return SimpleIDGenerator(prefix...)
	}

	LastIDs[hashString] = hashString
	return hashString
}


var LastIDs = make(map[string]string)


func checkIfIDExist(id string) bool {
	_, ok := LastIDs[id]
	return ok
}