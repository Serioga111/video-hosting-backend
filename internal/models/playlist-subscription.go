package models

type PlaylistSubscription struct {
	UserId     string `gorm:"not null" json:"user_id"`
	PlaylistId string `gorm:"not null" json:"playlist_id"`
}
