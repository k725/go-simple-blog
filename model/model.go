package model

import "github.com/k725/go-simple-blog/util"

func SetupDB() {
	setupInitialUser()
}

func setupInitialUser() {
	p, err := util.PasswordHash("passw0rd")
	if err != nil {
		panic(err)
	}
	a := User{
		UserID:   "admin",
		Password: p,
		Name:     "あどみん",
	}

	c := GetConnection()
	c.NewRecord(a)
	c.Create(&a)
}
