package main

import (
	"fmt"
	"github.com/frianlh/pokedex-api/configs"
	"github.com/frianlh/pokedex-api/routers"
	"log"
)

func main() {
	// flags for the standard logger
	log.SetFlags(log.Lshortfile | log.Lmicroseconds)

	// configuration
	newConfig := configs.NewConfig()
	config, err := newConfig.Read()
	if err != nil {
		log.Println(err)
		return
	}

	// route
	r := routers.SetupRoute(config)
	err = r.Listen(fmt.Sprintf(":%d", config.PortApi))
	if err != nil {
		log.Println(err)
		return
	}
}
