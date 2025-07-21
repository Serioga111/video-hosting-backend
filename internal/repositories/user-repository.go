package repositories

import (
	"video-hosting-backend/internal/models"
	"video-hosting-backend/internal/services"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(user *models.User) (*models.UserDTO, error)
	GetUserById(id string) (*models.User, error)
	GetUserByEmail(email string) (*models.UserDTO, error)
	UpdateUser(user *models.User) (*models.UserDTO, error)
	DeleteUser(id string) error
	ListUsers() ([]models.UserDTO, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) CreateUser(user *models.User) (*models.UserDTO, error) {
	if user.Id == "" {
		user.Id = uuid.NewString()
	}
	if err := r.db.Create(user).Error; err != nil {
		return nil, err
	}
	return services.ToUserDTO(user), nil
}

func (r *userRepository) GetUserById(id string) (*models.User, error) {
	var user *models.User
	if err := r.db.Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r *userRepository) GetUserByEmail(email string) (*models.UserDTO, error) {
	var user *models.User
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return services.ToUserDTO(user), nil
}

func (r *userRepository) UpdateUser(user *models.User) (*models.UserDTO, error) {
	if err := r.db.Save(user).Error; err != nil {
		return nil, err
	}
	return services.ToUserDTO(user), nil
}

func (r *userRepository) DeleteUser(id string) error {
	return r.db.Delete(&models.User{}, id).Error
}

func (r *userRepository) ListUsers() ([]models.UserDTO, error) {
	var users []models.UserDTO
	if err := r.db.Model(&models.User{}).
		Select("id", "username", "email", "created_at").
		Scan(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}
