package models

type CommentLike struct {
	CommentId string `gorm:"not null" json:"comment_id"`
	UserId    string `gorm:"not null" json:"user_id"`
	IsLike    bool   `gorm:"not null" json:"like"`
}
