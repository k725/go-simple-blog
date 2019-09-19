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
	CategoryID   int    `gorm:"column:categoryId"`
	UserName     string `gorm:"column:userName"`
	UserID       string `gorm:"column:userId"`
}

func fullArticleQueryBuilder() *gorm.DB {
	sel := "`articles`.*,"
	sel += "categories.name as categoryName, categories.id as categoryId,"
	sel += "users.name as userName, users.id as userId"

	return GetConnection().
		Table("articles").
		Select(sel).
		Joins("LEFT JOIN categories ON articles.category = categories.id").
		Joins("LEFT JOIN users ON articles.author = users.id").
		Order("created_at desc")
}

// GetArticle ...
func GetArticle(id int) ArticleFull {
	var a ArticleFull
	fullArticleQueryBuilder().
		Where("articles.id = ?", id).
		First(&a)
	return a
}

// GetAllArticles ...
// @todo fix
func GetAllArticles() []ArticleFull {
	var a []ArticleFull
	fullArticleQueryBuilder().
		Where("articles.deleted_at IS NULL").
		Scan(&a)
	return a
}

// GetArticlesByCategory ...
func GetArticlesByCategory(id int) []ArticleFull {
	var a []ArticleFull
	fullArticleQueryBuilder().
		Where("category = ?", id).
		Where("articles.deleted_at IS NULL").
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

func DeleteArticle(id int) error {
	c := GetConnection()
	return c.Where("id = ?", id).Delete(&Article{}).Error
}