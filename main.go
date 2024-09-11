package main

import (
	"context"
	"shorted/logger"
)

func main() {

	log := logger.New(context.Background())
	log.Info("Hello Shorted ..")
}
