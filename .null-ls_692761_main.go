package main

import (
	"github.com/C4ndid0/gojobsearh/config"
	"github.com/C4ndid0/gojobsearh/router"
)

var (
  logger *config.Logger
)

func main() {
  logger = config.GetLogger("main")
	//Initialize config
  err := config.Init()
  if err != nil {
    logger.Errorf("config initialization error: %v", err)
    return
  }

  //Initialize route
	router.Initializer()

} 
