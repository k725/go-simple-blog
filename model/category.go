package model

type Category struct {
	// gorm.Model
	Name string `gorm:"not null"`
}
