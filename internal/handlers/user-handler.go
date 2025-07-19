package handlers

import (
	"strconv"
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
	idUint, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid user ID"})
		return
	}
	user, err := h.repo.GetUserById(uint(idUint))
	if err != nil {
		c.JSON(404, gin.H{"error": "user not found"})
		return
	}

	c.JSON(200, gin.H{"user": user})

}

func (h *UserHandler) UpdateUser(c *gin.Context) {
	id := c.Param("id")
	idUint, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid user ID"})
		return
	}

	var input models.User
	if err := c.ShouldBindBodyWithJSON(&input); err != nil {
		c.JSON(400, gin.H{"error": "invald input"})
		return
	}
	input.Id = uint(idUint)

	user, err := h.repo.UpdateUser(&input)
	if err != nil {
		c.JSON(500, gin.H{"error": "coould not update user"})
		return
	}

	c.JSON(200, gin.H{"user": user})
}

func (h *UserHandler) DeleteUser(c *gin.Context) {
	id := c.Param("id")
	idUint, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid user ID"})
		return
	}
	if err := h.repo.DeleteUser(uint(idUint)); err != nil {
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
	if err != nil || bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)) != nil {
		c.JSON(401, gin.H{"error": "wrong email or password"})
		return
	}

	tokenString, err := services.GenerateRandomToken()
	if err != nil {
		c.JSON(500, gin.H{"error": "could not generate token"})
		return
	}

	token := models.Token{
		UserID:    user.Id,
		Token:     tokenString,
		ExpiresAt: time.Now().Add(24 * time.Hour),
	}

	if err = h.token.SaveToken(&token); err != nil {
		c.JSON(500, gin.H{"error": "could not save token"})
		return
	}

	c.JSON(200, gin.H{"token": tokenString})
}
