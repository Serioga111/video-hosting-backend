package models

type PlaybackHistory struct {
	PlaybackID string `gorm:"type:uuid" json:"playback_id"`
	VideoID    string `gorm:"not null" json:"video_id"`
	UserID     string `gorm:"not null" json:"user_id"`
	PlayedAt   string `gorm:"not null" json:"played_at"`
}
