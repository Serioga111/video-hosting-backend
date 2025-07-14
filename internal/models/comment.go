package models

type Comment struct {
	Id              uint   `gorm:"primaryKey" json:"id"`
	VideoId         uint   `gorm:"not null" json:"video_id"`
	UserId          uint   `gorm:"not null" json:"user_id"`                  // User who posted the comment
	ParentCommentId uint   `gorm:"default:0" json:"parent_comment_id"`       // ID of the parent comment for threaded comments
	Content         string `gorm:"type:text;not null" json:"content"`        // Content of the comment
	CreatedAt       string `gorm:"type:datetime;not null" json:"created_at"` // Timestamp of when the comment was created
}
