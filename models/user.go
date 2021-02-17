package models

import "github.com/jinzhu/gorm"

type Users struct {
	gorm.Model
	Login     string `gorm:"size:30;" json:"login"`
	Password  string `gorm:"size:30;" json:"password"`
	FirstName string `gorm:"size:30;" json:"first_name"`
	LastName  string `gorm:"size:30;" json:"last_name"`
}
