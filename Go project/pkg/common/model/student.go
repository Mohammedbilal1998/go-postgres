package model

import "gorm.io/gorm"

type Student struct {
	gorm.Model
	// ID uint  `json:"id" gorm:"column id"`
	Name string  `json:"name" gorm:"column name"`
	Phone string `json:"phone" gorm:"column phone"` 
	Email string `json:"email" gorm:"foreignkey;column email"`
}

