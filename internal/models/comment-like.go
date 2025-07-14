package models

type CommentLike struct {
	Id        uint `gorm:"primaryKey" json:"id"`
	CommentId uint `gorm:"not null" json:"comment_id"`
	UserId    uint `gorm:"not null" json:"user_id"`
	IsLike    bool `gorm:"not null" json:"like"`
}
