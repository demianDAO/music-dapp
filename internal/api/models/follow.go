package models

import "gorm.io/gorm"

type Follow struct {
	gorm.Model
	FollowerID string `json:"follower_id"`
	FolloweeID string `json:"followee_id"`
}
