package repositories

import (
	"time"
	"video-hosting-backend/internal/models"

	"gorm.io/gorm"
)

type UserDTO struct {
	Id        uint      `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

type UserRepository interface {
	CreateUser(user *models.User) (*UserDTO, error)
	GetUserById(id uint) (*UserDTO, error)
	UpdateUser(user *models.User) (*UserDTO, error)
	DeleteUser(id uint) error
	ListUsers() ([]UserDTO, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func toUserDTO(user *models.User) *UserDTO {
	return &UserDTO{
		Id:        user.Id,
		Username:  user.Username,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
	}
}

func (r *userRepository) CreateUser(user *models.User) (*UserDTO, error) {
	if err := r.db.Create(user).Error; err != nil {
		return nil, err
	}
	return toUserDTO(user), nil
}

func (r *userRepository) GetUserById(id uint) (*UserDTO, error) {
	var user models.User
	if err := r.db.First(&user, id).Error; err != nil {
		return nil, err
	}
	return toUserDTO(&user), nil
}

func (r *userRepository) UpdateUser(user *models.User) (*UserDTO, error) {
	if err := r.db.Save(user).Error; err != nil {
		return nil, err
	}
	return toUserDTO(user), nil
}

func (r *userRepository) DeleteUser(id uint) error {
	return r.db.Delete(&models.User{}, id).Error
}

func (r *userRepository) ListUsers() ([]UserDTO, error) {
	var users []UserDTO
	if err := r.db.Model(&models.User{}).
		Select("id", "username", "email", "created_at").
		Scan(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}
