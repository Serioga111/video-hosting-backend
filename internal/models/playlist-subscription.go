package models

type PlaylistSubscription struct {
	Id         uint `gorm:"primaryKey" json:"id"`
	UserId     uint `gorm:"not null" json:"user_id"`     // User who subscribed to the playlist
	PlaylistId uint `gorm:"not null" json:"playlist_id"` // ID of the playlist being subscribed to
}
