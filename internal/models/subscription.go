package models

type Subscription struct {
	Id        uint `gorm:"primaryKey" json:"id"`
	UserId    uint `gorm:"not null" json:"user_id"`
	ChannelId uint `gorm:"not null" json:"channel_id"`
}
