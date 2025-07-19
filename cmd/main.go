package main

import (
	"os"

	"video-hosting-backend/internal/config"
	"video-hosting-backend/internal/database"
	"video-hosting-backend/internal/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadEnvVariables()
	database.InitDB()

	router := gin.Default()

	//Swagger documentation
	router.StaticFile("/swagger.yaml", "./docs/swagger.yaml")
	router.StaticFile("/api-docs", "./docs/index.html")

	routes.SetupRoutes(router)

	router.Run(os.Getenv("PORT"))
}
