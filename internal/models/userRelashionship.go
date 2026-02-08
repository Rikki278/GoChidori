package models

import "time"

type UserRelationship struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`

	FollowerID uint         `gorm:"not null;index:idx_follower_following" json:"follower_id"`
	Follower   *UserProfile `gorm:"foreignKey:FollowerID;constraint:OnDelete:CASCADE" json:"follower,omitempty"`

	FollowingID uint         `gorm:"not null;index:idx_follower_following" json:"following_id"`
	Following   *UserProfile `gorm:"foreignKey:FollowingID;constraint:OnDelete:CASCADE" json:"following,omitempty"`
	
}

func (UserRelationship) TableName() string {
	return "user_relationships"
}