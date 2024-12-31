package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/priyanshuthapliyal/go-url-shortner/controllers"
)

func Routes(router *gin.Engine) {

	router.GET("/", controllers.Home)

	router.GET("/:url", controllers.Redirect)

	routes := router.Group("/api/v1")
	{
		routes.POST("/url", controllers.HandleNewUrl)
	}
}
