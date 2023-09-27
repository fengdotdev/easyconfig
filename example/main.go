package main

import (
	"easyconfig/easyconfig"
	"log"
	"os"
)

func main() {
	easyconfig.Dummy()


	stat, err := os.Stat("config.json")
	if err != nil {
		log.Println("no exist")
	}
	log.Println("exist", stat.Name())
}