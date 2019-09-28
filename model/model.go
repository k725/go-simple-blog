package model

import (
	"github.com/jinzhu/gorm"
	"github.com/k725/go-simple-blog/util"
)

func SetupDB() {
	c := GetConnection()

	setupInitialUser(c)
	setupCategory(c)
}

func setupCategory(con *gorm.DB) {
	var count int
	con.Model(Category{}).Count(&count)
	if count == 0 {
		a := Category{
			Name: "未設定",
		}
		con.NewRecord(a)
		con.Create(&a)
	}
}

func setupInitialUser(con *gorm.DB) {
	var count int
	con.Model(User{}).Count(&count)
	if count == 0 {
		p, err := util.PasswordHash("passw0rd")
		if err != nil {
			panic(err)
		}
		a := User{
			UserID:   "admin",
			Password: p,
			Name:     "あどみん",
		}
		con.NewRecord(a)
		con.Create(&a)
	}
}
