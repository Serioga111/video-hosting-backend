package models

type PlaylistToVideo struct {
	PlaylistId string `gorm:"not null" json:"playlist_id"`
	VideoId    string `gorm:"not null" json:"video_id"`
}
