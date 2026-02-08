package models

import (
	"time"

	"gorm.io/gorm"
)

type PostComment struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-"` // Soft delete for comments

	Content string `gorm:"type:text;not null" json:"content" validate:"required,max=500"`

	// Author of the comment
	UserID uint         `gorm:"not null;index" json:"user_id"`
	User   *UserProfile `gorm:"constraint:OnDelete:CASCADE" json:"user,omitempty"`

	// Post where the comment was made
	PostID uint           `gorm:"not null;index" json:"post_id"`
	Post   *CharacterPost `gorm:"constraint:OnDelete:CASCADE" json:"post,omitempty"`
}

func (PostComment) TableName() string {
	return "post_comments"
}

// BeforeUpdate GORM hook to set update timestamp
func (c *PostComment) BeforeUpdate(tx *gorm.DB) error {
	c.UpdatedAt = time.Now()
	return nil
}
