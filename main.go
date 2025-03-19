package main

import (
	"gopportunities/config"
	"gopportunities/router"
)

var (
	logger *config.Logger
)

func main() {
	// Initialize the logger
	logger = config.GetLogger("main")

	// Initialize the database connection
	err := config.Init()
	if err != nil {
		logger.Errorf("config init error: %s", err)
		return
	}

	// Initialize the router
	router.Initialize()
}
