package models

import "github.com/jinzhu/gorm"

type Comments struct {
	gorm.Model
	Name   string `gorm:"size:30;" json:"name"`
	Email  string `gorm:"not null" json:"email"`
	Body   string `gorm:"not null" json:"body"`
	PostId int    `gorm:"not null" json:"post_id"`
}
