package models

type PlaylistToVideo struct {
	Id         uint `gorm:"primaryKey" json:"id"`
	PlaylistId uint `gorm:"not null" json:"playlist_id"`
	VideoId    uint `gorm:"not null" json:"video_id"`
}
