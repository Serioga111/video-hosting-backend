package models

type PlaylistSubscription struct {
	Id         uint `gorm:"primaryKey" json:"id"`
	UserId     uint `gorm:"not null" json:"user_id"`
	PlaylistId uint `gorm:"not null" json:"playlist_id"`
}
