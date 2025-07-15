package routes

import (
	"video-hosting-backend/internal/database"
	"video-hosting-backend/internal/handlers"
	"video-hosting-backend/internal/repositories"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(rg *gin.RouterGroup) {
	UserRepository := repositories.NewUserRepository(database.DB)
	UserHandler := handlers.NewUserHandler(UserRepository)
	users := rg.Group("/users")
	users.POST("/", UserHandler.Register)
	users.GET("/", UserHandler.ListUsers)
	users.GET("/:id", UserHandler.GetUserById)
	users.PUT("/:id", UserHandler.UpdateUser)
	users.DELETE("/:id", UserHandler.DeleteUser)
}
