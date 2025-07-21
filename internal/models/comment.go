package models

import "time"

type Comment struct {
	Id              string    `gorm:"type:uuid" json:"id"`
	VideoId         string    `gorm:"not null" json:"video_id"`
	UserId          string    `gorm:"not null" json:"user_id"`
	ParentCommentId string    `gorm:"default:0" json:"parent_comment_id"`
	Content         string    `gorm:"type:text;not null" json:"content"`
	CreatedAt       time.Time `gorm:"not null" json:"created_at"`
}
