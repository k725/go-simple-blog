package model

import "github.com/jinzhu/gorm"

type Article struct {
	gorm.Model
	Title      string `gorm:"not null" sql:"type:text;"`
	Body       string `gorm:"not null" sql:"type:text;"`
	CategoryID uint   `gorm:"not null"`
	UserID     uint   `gorm:"not null"`
}

type ArticleFull struct {
	Article
	CategoryName string `gorm:"column:categoryName"`
	CategoryID   uint   `gorm:"column:categoryId"`
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
		Joins("LEFT JOIN categories ON articles.category_id = categories.id").
		Joins("LEFT JOIN users ON articles.user_id = users.id").
		Order("created_at desc")
}

// GetArticle ...
func GetArticle(id int) (ArticleFull, bool) {
	var a ArticleFull
	r := fullArticleQueryBuilder().
		Where("articles.id = ?", id).
		First(&a)

	return a, !r.RecordNotFound()
}

func GetArticles(offset, limit int) []ArticleFull {
	var a []ArticleFull
	fullArticleQueryBuilder().
		Where("articles.deleted_at IS NULL").
		Offset(offset).
		Limit(limit).
		Scan(&a)
	return a
}

func GetArticlesCount() int {
	var c int
	fullArticleQueryBuilder().
		Model(Article{}).
		Count(&c)
	return c
}

// GetArticlesByCategory ...
func GetArticlesByCategory(id, offset, limit int) []ArticleFull {
	var a []ArticleFull
	fullArticleQueryBuilder().
		Where("category_id = ?", id).
		Where("articles.deleted_at IS NULL").
		Offset(offset).
		Limit(limit).
		Scan(&a)
	return a
}

func GetArticlesByCategoryCount(id int) int {
	var c int
	fullArticleQueryBuilder().
		Model(Article{}).
		Where("category = ?", id).
		Count(&c)
	return c
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
