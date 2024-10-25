package model

import "gorm.io/gorm"

type Course struct {
	gorm.Model
	// ID  uint  `json:"id" gorm:"column id"`
	Name string  `json:"name" gorm:"column name"`
	Student_ID uint `json:"studentid" gorm:"column studentid"`
}