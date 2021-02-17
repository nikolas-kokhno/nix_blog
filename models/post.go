package models

type Posts struct {
	ID     int64  `gorm:"primary_key;auto_increment;not_null" json:"id"`
	Title  string `gorm:"size:30;" json:"title"`
	Body   string `gorm:"not null" json:"body"`
	UserID int64  `gorm:"not null" json:"user_id"`
}
