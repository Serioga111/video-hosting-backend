package routes

import "github.com/gin-gonic/gin"

func SetupRoutes(router *gin.Engine) {
	api := router.Group("/api")

	RegisterUserRoutes(api)
}
