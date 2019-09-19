package model

import "github.com/jinzhu/gorm"

type Category struct {
	gorm.Model
	Name string `gorm:"not null"`
}

func GetAllCategories() []Category {
	var c []Category
	GetConnection().Find(&c)
	return c
}