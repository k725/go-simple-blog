package model

import "github.com/jinzhu/gorm"

type Article struct {
	gorm.Model
	Title    string `gorm:"not null"`
	Body     string `gorm:"not null"`
	Category int    `gorm:"not null"`
	Author   int    `gorm:"not null"`
}

// GetArticle ...
func GetArticle(id int) Article {
	var a Article
	GetConnection().Where("id = ?", id).First(&a)
	return a
}

// GetAllArticles ...
// @todo fix
func GetAllArticles() []Article {
	var a []Article
	GetConnection().Find(&a)
	return a
}
