package models

type Token struct {
	ID        string `gorm:"primaryKey" json:"id"`
	UserID    string `gorm:"not null" json:"user_id"`
	Token     string `gorm:"type:varchar(255);not null;unique" json:"token"`
	IssuedAt  string `json:"issued_at" gorm:"not null"`
	ExpiresAt string `json:"expires_at" gorm:"not null"`
}
