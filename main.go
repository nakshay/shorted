package main

import (
	"context"
	"shorted/logger"
)

var db = make(map[string]string)

func main() {
	r := setupRouter()
	err := r.Run(":8080")
	log := logger.New(context.Background())
	if err != nil {
		log.Error("Fatal error while starting server")
	}

}
