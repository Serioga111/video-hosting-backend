package models

type VideoLike struct {
	VideoId string `gorm:"not null" json:"video_id"`
	UserId  string `gorm:"not null" json:"user_id"`
	IsLike  bool   `gorm:"not null" json:"like"`
}
