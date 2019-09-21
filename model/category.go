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

func HasCategory(id int) bool {
	var c = 0
	GetConnection().
		Model(&Category{}).
		Where("id = ?", id).
		Count(&c)
	return c != 0
}

func InsertCategory(a Category) error {
	c := GetConnection()
	c.NewRecord(a)
	return c.Create(&a).Error
}

func DeleteCategory(id int) error {
	c := GetConnection()
	return c.Where("id = ?", id).Delete(&Category{}).Error
}
