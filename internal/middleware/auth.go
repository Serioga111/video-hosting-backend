package middleware

import (
	"video-hosting-backend/internal/services"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHead := c.GetHeader("Authorization")
		if authHead == "" {
			c.AbortWithStatusJSON(401, gin.H{"error": "Unauthorized"})
			return
		}
		tokenStr, err := services.SplitBearerToken(authHead)
		if err != nil {
			c.AbortWithStatusJSON(401, gin.H{"error": "Invalid authorization header format"})
			return
		}

		claims, err := services.ValidateAccessToken(tokenStr)
		if err != nil {
			c.AbortWithStatusJSON(401, gin.H{"error": "Invalid or expired token"})
			return
		}

		c.Set("userId", claims.UserID)
		c.Set("rawToken", tokenStr)

		c.Next()
	}
}
