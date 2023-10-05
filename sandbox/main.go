package main

import (
	"easyconfig/easyconfig/disk"
	"fmt"
)

//PENDING
//IMPORTANT



func main (){
  
	b:= []byte("hello world")

	d := disk.NewDisk("/tmp")

	err := d.Write(b, "tmp", "test.txt")
	if err != nil {
		fmt.Println(err.V())
	}
}