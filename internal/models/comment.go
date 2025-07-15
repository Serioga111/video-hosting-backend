package models

import "time"

type Comment struct {
	Id              uint      `gorm:"primaryKey" json:"id"`
	VideoId         uint      `gorm:"not null" json:"video_id"`
	UserId          uint      `gorm:"not null" json:"user_id"`
	ParentCommentId uint      `gorm:"default:0" json:"parent_comment_id"`
	Content         string    `gorm:"type:text;not null" json:"content"`
	CreatedAt       time.Time `gorm:"not null" json:"created_at"`
}
