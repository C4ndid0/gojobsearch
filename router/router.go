package router

import "github.com/gin-gonic/gin"

func Initializer() {
	//Initialize Router
	router := gin.Default()

	//Initialize Routes
	initializeRoutes(router)

	//Get port from the environment

	router.Run(":8080")
}

//Run the server
