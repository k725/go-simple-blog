package model

func setupDB() {
	a := User{
		UserID:   "admin",
		Password: "",
		Name:     "あどみん",
	}

	c := GetConnection()
	c.NewRecord(a)
	c.Create(&a)

}
