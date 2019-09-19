package model

import "github.com/jinzhu/gorm"

type Article struct {
	gorm.Model
	Title    string `gorm:"not null" sql:"type:text;"`
	Body     string `gorm:"not null" sql:"type:text;"`
	Category int    `gorm:"not null"`
	Author   int    `gorm:"not null"`
}

type ArticleFull struct {
	Article
	CategoryName string `gorm:"column:categoryName"`
	CategoryID int `gorm:"column:categoryId"`
}

// GetArticle ...
func GetArticle(id int) ArticleFull {
	var a ArticleFull
	GetConnection().
		Table("articles").
		Select("`articles`.*, categories.name as categoryName, categories.id as categoryId").
		Joins("LEFT JOIN categories ON articles.category = categories.id").
		Order("created_at desc").
		Where("articles.id = ?", id).
		First(&a)
	return a
}

// GetAllArticles ...
// @todo fix
func GetAllArticles() []ArticleFull {
	var a []ArticleFull
	GetConnection().
		Table("articles").
		Select("`articles`.*, categories.name as categoryName").
		Joins("LEFT JOIN categories ON articles.category = categories.id").
		Order("created_at desc").
		Scan(&a)
	return a
}

// GetArticlesByCategory ...
func GetArticlesByCategory(id int) []ArticleFull {
	var a []ArticleFull
	GetConnection().
		Table("articles").
		Select("`articles`.*, categories.name as categoryName").
		Joins("LEFT JOIN categories ON articles.category = categories.id").
		Order("created_at desc").
		Where("category = ?", id).
		Scan(&a)
	return a
}

func InsertArticle(a Article) error {
	c := GetConnection()
	c.NewRecord(a)
	return c.Create(&a).Error
}

func UpdateArticle(a Article) error {
	c := GetConnection()
	return c.Model(&a).Updates(a).Error
}
