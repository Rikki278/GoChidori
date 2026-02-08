package models

import (
	"time"
)

type PostLike struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time `json:"created_at"`

	// User who liked the post
	UserID uint         `gorm:"not null;index:idx_user_post_like" json:"user_id"`
	User   *UserProfile `gorm:"constraint:OnDelete:CASCADE" json:"user,omitempty"`

	// Post that was liked
	PostID uint           `gorm:"not null;index:idx_user_post_like" json:"post_id"`
	Post   *CharacterPost `gorm:"constraint:OnDelete:CASCADE" json:"post,omitempty"`
}

func (PostLike) TableName() string {
	return "post_likes"
}
