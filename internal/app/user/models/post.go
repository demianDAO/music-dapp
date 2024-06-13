package models

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	UserAddr string `json:"user_addr"`
	Content  string `json:"content"`
}
