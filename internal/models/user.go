package models

import (
	"time"
)

type User struct {
	Id        string    `gorm:"type:uuid;primaryKey" json:"id"`
	Username  string    `gorm:"type:varchar(50);unique;not null"`
	Email     string    `gorm:"type:varchar(100);unique;not null"`
	Password  string    `gorm:"type:varchar(255);not null"`
	Settings  string    `gorm:"type:text;not null" json:"settings"` // JSON string to store user settings
	CreatedAt time.Time `gorm:"not null"`
}

type UserDTO struct {
	Id        string    `json:"id" `
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}
