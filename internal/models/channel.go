package models

import "time"

type Channel struct {
	Id               uint       `gorm:"primaryKey" json:"id"`
	UserId           uint       `gorm:"not null" json:"user_id"`
	Name             string     `gorm:"type:varchar(100);not null" json:"name"`
	Description      string     `gorm:"type:text;not null" json:"description"`
	AvatarUrl        string     `gorm:"type:varchar(255);not null" json:"avatar_url"`
	TotalViews       int        `gorm:"default:0" json:"total_views"`
	CreatedAt        time.Time  `gorm:"not null" json:"created_at"`
	SubscribersCount int        `gorm:"default:0" json:"subscribers_count"`
	Subscribers      []User     `gorm:"many2many:channel_subscribers;" json:"subscribers"`
	Playlists        []Playlist `gorm:"foreignKey:ChannelId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"playlists"`
	Videos           []Video    `gorm:"foreignKey:ChannelId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"videos"`
}
