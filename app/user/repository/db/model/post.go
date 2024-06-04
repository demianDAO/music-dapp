package model

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	AuthorAddr string `json:"author_addr"`
	Content    string `json:"content"`
}
