package model

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	UserID   string `gorm:"not null"`
	Password string `gorm:"not null"`
	Name     string `gorm:"not null"`
}
