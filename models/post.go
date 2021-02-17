package models

import "github.com/jinzhu/gorm"

type Posts struct {
	gorm.Model
	Title  string `gorm:"size:30;" json:"title"`
	Body   string `gorm:"not null" json:"body"`
	UserID int    `gorm:"not null" json:"user_id"`
}
