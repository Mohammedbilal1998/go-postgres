package model

import (
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Title string `json:"title" gorm:"column:title"`
	Author string `json:"author" gorm:"column:author"`
	Description string `json:"description" gorm:"column:description"`
}