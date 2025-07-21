package handlers

import (
	"time"
	"video-hosting-backend/internal/models"
	"video-hosting-backend/internal/repositories"
	"video-hosting-backend/internal/services"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type UserHandler struct {
	repo  repositories.UserRepository
	token repositories.TokenRepository
}

func NewUserHandler(repo repositories.UserRepository, token repositories.TokenRepository) *UserHandler {
	return &UserHandler{
		repo:  repo,
		token: token,
	}
}

func (h *UserHandler) Register(c *gin.Context) {
	var input models.User
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": "invalid input"})
		return
	}

	hashedPassword, err := services.HashPassword(input.Password)
	if err != nil {
		c.JSON(500, gin.H{"error": "could not hash password"})
		return
	}

	input.Password = hashedPassword

	user, err := h.repo.CreateUser(&input)
	if err != nil {
		c.JSON(500, gin.H{"error": "could not create user"})
		return
	}

	c.JSON(201, gin.H{"user": user})
}

func (h *UserHandler) GetUserById(c *gin.Context) {
	id := c.Param("id")
	user, err := h.repo.GetUserById(id)
	if err != nil {
		c.JSON(404, gin.H{"error": "user not found"})
		return
	}

	c.JSON(200, gin.H{"user": user})

}

func (h *UserHandler) GetUserByEmail(c *gin.Context) {
	email := c.Param("email")
	user, err := h.repo.GetUserByEmail(email)
	if err != nil {
		c.JSON(404, gin.H{"error": "user not found"})
		return
	}

	c.JSON(200, gin.H{"user": user})

}

func (h *UserHandler) UpdateUser(c *gin.Context) {
	id := c.Param("id")

	var input models.User
	if err := c.ShouldBindBodyWithJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": "invald input"})
		return
	}
	input.Id = id

	user, err := h.repo.UpdateUser(&input)
	if err != nil {
		c.JSON(500, gin.H{"error": "coould not update user"})
		return
	}

	c.JSON(200, gin.H{"user": user})
}

func (h *UserHandler) DeleteUser(c *gin.Context) {
	id := c.Param("id")
	if err := h.repo.DeleteUser(id); err != nil {
		c.JSON(500, gin.H{"error": "could not delete user"})
		return
	}

	c.Status(204)
}

func (h *UserHandler) ListUsers(c *gin.Context) {
	users, err := h.repo.ListUsers()
	if err != nil {
		c.JSON(500, gin.H{"error": "could not list users"})
		return
	}

	c.JSON(200, gin.H{"users": users})
}

func (h *UserHandler) Login(c *gin.Context) {
	var input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": "invalid input"})
		return
	}

	user, err := h.repo.GetUserByEmail(input.Email)
	if err != nil {
		c.JSON(401, gin.H{"error": "wrong email"})
		return
	}

	userById, err := h.repo.GetUserById(user.Id)
	if err != nil || bcrypt.CompareHashAndPassword([]byte(userById.Password), []byte(input.Password)) != nil {
		c.JSON(401, gin.H{"error": "wrong password"})
		return
	}

	accessToken, err := services.GenerateAccessToken(string(user.Id))
	if err != nil {
		c.JSON(500, gin.H{"error": "could not generate access token"})
		return
	}

	tokenString, err := services.GenerateRandomToken()
	if err != nil {
		c.JSON(500, gin.H{"error": "could not generate token"})
		return
	}

	token := models.Token{
		ID:        tokenString,
		IssuedAt:  time.Now(),
		UserID:    user.Id,
		Token:     tokenString,
		ExpiresAt: time.Now().Add(30 * 24 * time.Hour),
	}

	if err = h.token.SaveToken(&token); err != nil {
		c.JSON(500, gin.H{"error": "could not save token"})
		return
	}

	c.SetCookie(
		"refresh_token",
		tokenString,
		30*24*60*60, // 30 days
		"/",
		"",
		true,
		true,
	)

	c.JSON(200, gin.H{"token": accessToken})
}

func (h *UserHandler) Logout(c *gin.Context) {
	refreshToken, err := c.Cookie("refresh_token")
	if err != nil {
		c.JSON(401, gin.H{"error": "refresh token not found"})
		return
	}

	if err := h.token.DeleteToken(refreshToken); err != nil {
		c.JSON(500, gin.H{"error": "could not delete token"})
		return
	}

	c.Status(204)
}

func (h *UserHandler) RefreshToken(c *gin.Context) {
	refreshToken, err := c.Cookie("refresh_token")
	if err != nil {
		c.JSON(401, gin.H{"error": "refresh token not found"})
		return
	}

	storedToken, err := h.token.GetValidToken(refreshToken)
	if err != nil {
		c.JSON(401, gin.H{"error": "invalid or expired refresh token"})
		return
	}

	accessToken, err := services.GenerateAccessToken(storedToken.UserID)
	if err != nil {
		c.JSON(500, gin.H{"error": "could not generate access token"})
		return
	}

	c.JSON(200, gin.H{
		"access_token": accessToken,
	})
}
