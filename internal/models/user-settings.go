package models

type UserSettings struct {
	ID       uint   `gorm:"primaryKey" json:"id"`
	UserID   uint   `gorm:"not null;unique" json:"user_id"`
	Settings string `gorm:"type:text;not null" json:"settings"` // JSON string to store user settings
}
