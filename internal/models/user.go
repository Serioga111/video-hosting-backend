package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Id        uint   `gorm:"primaryKey"`
	Username  string `gorm:"type:varchar(50);unique;not null"`
	Email     string `gorm:"type:varchar(100);unique;not null"`
	Password  string `gorm:"type:varchar(255);not null"`
	CreatedAt string `gorm:"type:datetime;not null"`
}
