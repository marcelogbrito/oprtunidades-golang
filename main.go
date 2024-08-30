package main

import (
	"github.com/marcelogbrito/oprtunidades-golang/config"
	"github.com/marcelogbrito/oprtunidades-golang/router"
)

var (
	logger *config.Logger
)

func main() {

	logger = config.GetLogger("main")
	// Initialize configs
	err := config.Init()
	if err != nil {
		logger.Errorf("config initialization error: %v", err)
		return
	}

	//inicia o router
	router.Initialize()
}
