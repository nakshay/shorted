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
	logger := loggingUtil.GetLogger(context.Background())
	logger.Info("Starting shorted service")
	r := setupRouter(config)
	err = r.Run(":8080")

	if err != nil {
		logger.Error("Fatal error while starting server")
	}

}
