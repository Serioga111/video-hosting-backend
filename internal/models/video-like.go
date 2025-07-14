package models

type VideoLike struct {
	Id      uint `gorm:"primaryKey" json:"id"`
	VideoId uint `gorm:"not null" json:"video_id"`
	UserId  uint `gorm:"not null" json:"user_id"`
	IsLike  bool `gorm:"not null" json:"like"`
}
