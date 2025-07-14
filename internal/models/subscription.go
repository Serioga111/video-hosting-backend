package models

type Subscription struct {
	Id        uint `gorm:"primaryKey" json:"id"`
	UserId    uint `gorm:"not null" json:"user_id"`    // User who subscribed to the channel
	ChannelId uint `gorm:"not null" json:"channel_id"` // ID of the channel being subscribed to
}
