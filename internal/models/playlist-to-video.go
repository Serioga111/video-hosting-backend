package models

type PlaylistToVideo struct {
	Id         uint `gorm:"primaryKey" json:"id"`
	PlaylistId uint `gorm:"not null" json:"playlist_id"` // ID of the playlist
	VideoId    uint `gorm:"not null" json:"video_id"`    // ID of the video in the playlist
}
