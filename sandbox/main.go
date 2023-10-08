package main

import (
	"fmt"

	"github.com/fengdotdev/easyconfig/easyconfig/disk"
)

//PENDING
//IMPORTANT



func main (){
  
	b:= []byte("hello world")
	d, err := disk.NewStorage("/tmp")
	if err != nil {
		panic(err)
	}
	fmt.Println(d, b)// for debugging

}