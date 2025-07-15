package handlers

import (
	"strconv"
	"video-hosting-backend/internal/models"
	"video-hosting-backend/internal/repositories"
	"video-hosting-backend/internal/services"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	repo repositories.UserRepository
}

func NewUserHandler(repo repositories.UserRepository) *UserHandler {
	return &UserHandler{repo: repo}
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
