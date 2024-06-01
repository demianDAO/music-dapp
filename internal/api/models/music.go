package models

import "gorm.io/gorm"

type Music struct {
	gorm.Model
	ID       string `json:"id"`
	Title    string `json:"title"`
	Overview string `json:"overview"`
	Author   string `json:"author"`
}
