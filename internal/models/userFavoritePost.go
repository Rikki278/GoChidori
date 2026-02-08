package models

import "time"

type UserFavoritePost struct {
	ID uint `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time `json:"created_at"`

	UserID uint `gorm:"not null;index" json:"user_id"`
	User *UserProfile `gorm:"constraint:OnDelete:CASCADE" json:"user,omitempty"`
	
	PostID uint `gorm:"not null;index" json:"post_id"`
	Post *CharacterPost `gorm:"constraint:OnDelete:CASCADE" json:"post,omitempty"`
	
}

func (UserFavoritePost) TableName() string {
	return "user_favorite_posts"
}
