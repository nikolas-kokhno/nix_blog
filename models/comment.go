package models

type Comments struct {
	ID     int64  `gorm:"primary_key;auto_increment;not_null" json:"id"`
	Name   string `gorm:"size:30;" json:"name"`
	Email  string `gorm:"not null" json:"email"`
	Body   string `gorm:"not null" json:"body"`
	PostId int64  `gorm:"not null" json:"post_id"`
}
