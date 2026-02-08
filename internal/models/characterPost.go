package models

import (
	"time"

	"gorm.io/gorm"
)

type CharacterPost struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-"` // soft delete

	// main fields
	Title       string `gorm:"not null;size:100" json:"title" validate:"required,min=3,max=100"`
	Content     string `gorm:"not null;type:text" json:"content" validate:"required"`
	Description string `gorm:"size:255" json:"description,omitempty" validate:"max=255"`

	ImageUrl     string `gorm:"size:255" json:"image_url,omitempty" validate:"omitempty,url"`
	ThumbnailUrl string `gorm:"size:255" json:"thumbnail_url,omitempty" validate:"omitempty,url"`

	Tags     string `gorm:"size:500" json:"tags,omitempty"`
	Category string `gorm:"size:100" json:"category,omitempty"`

	ViewCount int `gorm:"default:0" json:"view_count"`
	LikeCount int `gorm:"default:0" json:"like_count"`

	// foreign keys
	UserID uint         `gorm:"not null;index" json:"user_id"`
	User   *UserProfile `gorm:"constraint:OnDelete:CASCADE" json:"user,omitempty"`

	// relationships
	FavoritedBy []UserFavoritePost `gorm:"foreignKey:PostID;constraint:OnDelete:CASCADE" json:"favorited_by,omitempty"`
	Likes       []PostLike         `gorm:"foreignKey:PostID;constraint:OnDelete:CASCADE" json:"likes,omitempty"`
	Comments    []PostComment      `gorm:"foreignKey:PostID;constraint:OnDelete:CASCADE" json:"comments,omitempty"`
}

func (CharacterPost) TableName() string {
	return "character_posts"
}

func (p *CharacterPost) BeforeCreate(tx *gorm.DB) (err error) {
	if p.Title == "" {
		return gorm.ErrInvalidData
	}

	return nil
}

func (p *CharacterPost) BeforeUpdate(tx *gorm.DB) (err error) {
	p.UpdatedAt = time.Now()
	return nil
}
