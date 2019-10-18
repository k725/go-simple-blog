package model

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	UserID   string `gorm:"size:16;unique_index;not null"`
	Password string `gorm:"not null"`
	Name     string `gorm:"not null"`
}

func GetUserByUserId(userId string) User {
	var u User
	GetConnection().
		Where("user_id = ?", userId).
		First(&u)
	return u
}

func UpdateUser(u User) error {
	c := GetConnection()
	return c.Model(&u).Updates(u).Error
}
