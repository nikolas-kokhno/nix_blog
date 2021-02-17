package models

type Users struct {
	ID       int64  `gorm:"primary_key;auto_increment;not_null" json:"id"`
	Name     string `gorm:"size:80" json:"name"`
	Username string `gorm:"size:80;unique" json:"username"`
	Password string `gorm:"size:80" json:"password"`
	Email    string `gorm:"size:60" json:"email"`
	Phone    string `gorm:"size:30" json:"phone"`
	Website  string `gorm:"size:40" json:"website"`
}
