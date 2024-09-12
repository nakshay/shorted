package main

import (
	"context"
	"shorted/configuration"
	"shorted/logger"
)

var db = make(map[string]string)

func main() {
	config, err := configuration.NewConfigLoader().LoadConfig("./configuration/config.json")
	if err != nil {
		panic("Error while reading configuration file")
	}

	r := setupRouter(config)
	err = r.Run(":8080")
	log := logger.New(context.Background())
	if err != nil {
		log.Error("Fatal error while starting server")
	}

}
