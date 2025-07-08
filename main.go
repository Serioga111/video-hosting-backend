package main

import (
	"video-hosting-backend/initialaizers"

	"github.com/gin-gonic/gin"
)

func init() {
	initialaizers.LoadEnvVariables()
}

func main() {
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	router.Run()
}
