package main

import (
	"log"

	"github.com/Mohammedbilal1998/go-postgres/pkg/common/config"
)

func main () {
	config, err := config.GetConfig()
	if err != nil {
		log.Fatalln("unable to read the config")
	}
	log.Println(config)
}