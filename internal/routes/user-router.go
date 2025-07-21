package routes

import (
	"video-hosting-backend/internal/database"
	"video-hosting-backend/internal/handlers"
	"video-hosting-backend/internal/middleware"
	"video-hosting-backend/internal/repositories"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(rg *gin.RouterGroup) {
	UserRepository := repositories.NewUserRepository(database.DB)
	TokenRepository := repositories.NewTokenRepository(database.DB)
	UserHandler := handlers.NewUserHandler(UserRepository, TokenRepository)

	rg.GET("/users", UserHandler.ListUsers)
	rg.POST("/register", UserHandler.Register)
	rg.POST("/login", UserHandler.Login)
	rg.POST("/auth/refresh", UserHandler.RefreshToken)

	users := rg.Group("/user")
	{
		users.GET("/:id", middleware.AuthMiddleware(), UserHandler.GetUserById)
		users.GET("/by-email/:email", middleware.AuthMiddleware(), UserHandler.GetUserByEmail)
		users.PUT("/:id", middleware.AuthMiddleware(), UserHandler.UpdateUser)
		users.DELETE("/:id", middleware.AuthMiddleware(), UserHandler.DeleteUser)
		users.DELETE("/logout", middleware.AuthMiddleware(), UserHandler.Logout)
	}
}
