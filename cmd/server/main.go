package main

import (
	"os"
	"video-hosting-backend/config"
	"video-hosting-backend/internal/database"
	"video-hosting-backend/internal/models"

	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadEnvVariables()
	database.InitDB()

	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		var user models.User
		if err := database.DB.First(&user).Error; err != nil {
			c.JSON(500, gin.H{"error": "could not fetch user"})
			return
		}

		c.JSON(200, user)
	})

	router.Run(os.Getenv("PORT"))
}
