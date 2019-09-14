package model

import "github.com/jinzhu/gorm"

type Article struct {
	gorm.Model
	Title    string `gorm:"not null"`
	Body     string `gorm:"not null"`
	Category int    `gorm:"not null"`
	Author   int    `gorm:"not null"`
}
