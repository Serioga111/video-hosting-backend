package models

type Subscription struct {
	UserId    string `gorm:"not null" json:"user_id"`
	ChannelId string `gorm:"not null" json:"channel_id"`
}
