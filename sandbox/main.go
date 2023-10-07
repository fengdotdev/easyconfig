package main

import (
	"easyconfig/easyconfig/disk"
	"fmt"
)

//PENDING
//IMPORTANT



func main (){
  
	b:= []byte("hello world")
	d, err := disk.NewDisk("/tmp")
	if err != nil {
		panic(err)
	}
	fmt.Println(d, b)// for debugging

}