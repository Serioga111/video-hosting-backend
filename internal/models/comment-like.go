package models

type CommentLike struct {
	Id        uint `gorm:"primaryKey" json:"id"`
	CommentId uint `gorm:"not null" json:"comment_id"`
	UserId    uint `gorm:"not null" json:"user_id"` // User who liked or disliked the comment
	IsLike    bool `gorm:"not null" json:"like"`    // true for like, false for dislike
}
