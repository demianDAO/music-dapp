package models

import "gorm.io/gorm"

type Song struct {
	gorm.Model
	Title      string `json:"title"`
	Overview   string `json:"overview"`
	AuthorAddr string `json:"author_addr"`
}
