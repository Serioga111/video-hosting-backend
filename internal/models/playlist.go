package models

import "time"

type Playlist struct {
	Id         uint      `gorm:"primaryKey" json:"id"`
	ChannelId  uint      `gorm:"not null" json:"channel_id"`
	Title      string    `gorm:"type:varchar(100);not null" json:"title"`
	PreviewUrl string    `gorm:"type:varchar(255);not null" json:"preview_url"`
	CreatedAt  time.Time `gorm:"not null" json:"created_at"`
	ViewsCount int       `gorm:"default:0" json:"views_count"`
}
