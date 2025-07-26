package main

import (
	"github.com/VictorMarri/Gopportunities.git/config"
	"github.com/VictorMarri/Gopportunities.git/router"
)

var (
	log config.Logger
)

func main() {

	logger := config.GetLogger("Initializing main method")

	//Initialize project configs
	err := config.Init()
	if err != nil {
		//panic(err) //Kills the app if configuration has any error
		logger.Errorf("Config initialization error: %v", err)
		return
	}

	//Initialize router
	router.Initialize()
}
