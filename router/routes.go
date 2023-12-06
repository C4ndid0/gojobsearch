package router

import (
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	docs "github.com/C4ndid0/gojobsearh/docs"
	"github.com/C4ndid0/gojobsearh/handler"
)

func initializeRoutes(router *gin.Engine) {
	// Initialize Handler
	handler.InitializeHandler()
	basePath := "/api/v1"
	docs.SwaggerInfo.BasePath = basePath
	v1 := router.Group(basePath)

	{

		v1.GET("/opening", handler.ShowOpeningHandler)
		v1.POST("/opening", handler.CreateOpeningHandler)
		v1.PUT("/opening", handler.UpdateOpeningHandler)
		v1.DELETE("/opening", handler.DeleteOpeningHandler)
		v1.GET("/openings", handler.ListOpeningHandler)

	}

	// Initialize swaggo
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
