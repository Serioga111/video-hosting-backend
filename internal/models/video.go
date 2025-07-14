package models

import (
	"time"
)

type Video struct {
	Id             uint      `gorm:"primaryKey" json:"id"`
	ChannelId      uint      `gorm:"not null" json:"channel_id"`
	Title          string    `gorm:"type:varchar(255);not null" json:"title"`
	Description    string    `gorm:"type:text;not null" json:"description"`
	PreviewUrl     string    `gorm:"type:varchar(255);not null" json:"preview_url"`
	VideoUrl       string    `gorm:"type:varchar(255);not null" json:"video_url"`
	ReleaseDate    time.Time `gorm:"not null" json:"release_date"`
	IsPrivate      bool      `gorm:"default:false" json:"is_private"`
	LikesCount     int       `gorm:"default:0" json:"likes_count"`
	DislikesCount  int       `gorm:"default:0" json:"dislikes_count"`
	CommentrsCount int       `gorm:"default:0" json:"comments_count"`
	Comments       []Comment `gorm:"foreignKey:VideoId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"comments"`
}
