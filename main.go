package main

import (
	"context"
	"shorted/configuration"
	"shorted/loggingUtil"
)

func main() {
	config, err := configuration.NewConfigLoader().LoadConfig("./configuration/config.json")
	if err != nil {
		panic("Error while reading configuration file")
	}
	log := loggingUtil.GetLogger(context.Background())
	r := setupRouter(config)
	err = r.Run(":8080")

	if err != nil {
		log.Error("Fatal error while starting server")
	}

}
