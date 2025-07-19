package repositories

import (
	"video-hosting-backend/internal/models"

	"gorm.io/gorm"
)

type TokenRepository interface {
	SaveToken(token *models.Token) error
	GetValidToken(tokenString string) (*models.Token, error)
	DeleteToken(tokenString string) error
}

type tokenRepository struct {
	db *gorm.DB
}

func NewTokenRepository(db *gorm.DB) TokenRepository {
	return &tokenRepository{db: db}
}

func (r *tokenRepository) SaveToken(token *models.Token) error {
	if err := r.db.Create(token).Error; err != nil {
		return err
	}
	return nil
}

func (r *tokenRepository) GetValidToken(tokenString string) (*models.Token, error) {
	var token models.Token
	if err := r.db.Where("token = ?", tokenString).First(&token).Error; err != nil {
		return nil, err
	}
	return &token, nil
}

func (r *tokenRepository) DeleteToken(tokenString string) error {
	var token models.Token
	if err := r.db.Where("token = ?", tokenString).First(&token).Error; err != nil {
		return err
	}
	if err := r.db.Delete(&token).Error; err != nil {
		return err
	}
	return nil
}
