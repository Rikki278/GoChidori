package models

import (
	"time"

	"gorm.io/gorm"
)

type UserProfile struct {
	ID        uint       `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"-"` // soft delete

	// main fields
	Username string `gorm:"not null;size:50" json:"username" validate:"required,min=3,max=50"`
	Email    string `gorm:"uniqueIndex;not null;size:100" json:"email" validate:"required,email"`
	Password string `gorm:"not null" json:"-" validate:"required,min=8"` // not shown in json

	Role UserRole `gorm:"type:varchar(20);not null;default:'ROLE_USER'" json:"role"`

	// profile info
	FirstName       string     `gorm:"size:50" json:"first_name,omitempty" validate:"max=50"`
	LastName        string     `gorm:"size:50" json:"last_name,omitempty" validate:"max=50"`
	Bio             string     `gorm:"size:255" json:"bio,omitempty" validate:"max=500"`
	ProfileImageUrl string     `gorm:"size:255" json:"profile_image_url,omitempty" validate:"omitempty,url"`
	LastLogin       *time.Time `gorm:"column:last_login" json:"last_login,omitempty"`

	RefreshToken string `gorm:"size:255" json:"refresh_token,omitempty" validate:"omitempty,max=255"`

	// relationships one to many
	CharacterPosts []CharacterPost    `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE" json:"character_posts,omitempty"`
	FavoritePosts  []UserFavoritePost `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE" json:"favorite_posts,omitempty"`
	Likes          []PostLike         `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE" json:"likes,omitempty"`
	Comments       []PostComment      `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE" json:"comments,omitempty"`
	Following      []UserRelationship `gorm:"foreignKey:FollowerID;constraint:OnDelete:CASCADE" json:"following,omitempty"`
	Followers      []UserRelationship `gorm:"foreignKey:FollowingID;constraint:OnDelete:CASCADE" json:"followers,omitempty"`
}

func (UserProfile) TableName() string {
	return "user_profiles"
}

func (u *UserProfile) BeforeCreate(tx *gorm.DB) error {
	if u.Role == "" {
		u.Role = RoleUser
	}
	return nil
}

func (u *UserProfile) BeforeUpdate(tx *gorm.DB) error {
	u.UpdatedAt = time.Now()
	return nil
}
