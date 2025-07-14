package database

import (
	"fmt"
	"log"
	"os"
	"video-hosting-backend/internal/models"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
		os.Getenv("SSL_MODE"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	err = db.AutoMigrate(
		&models.User{},
		&models.Channel{},
		&models.Video{},
		&models.Playlist{},
		&models.Comment{},
		&models.VideoLike{},
		&models.CommentLike{},
		&models.PlaylistToVideo{},
		&models.Subscription{},
		&models.PlaylistSubscription{},
		&models.PlaybackHistory{},
		&models.Token{},
		&models.UserSettings{},
	)

	if err != nil {
		log.Fatalf("Failed to automigrate database: %v", err)
	}

	DB = db
}
