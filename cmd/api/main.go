package main

import (
	"log"

	"github.com/docker/docker/daemon/config"
)

func main() {
	//Define config
	log.Println("Definisi Config.....")
	config, err := config.New("prod", ".")
	if err != nil {
		log.Printf("error loading config -> %v\n", err)
	}

	//Define Logger
	//Define Database
	//Define Server Api
	//Start The Server

}
