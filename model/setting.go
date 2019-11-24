package model

import "github.com/jinzhu/gorm"

type Setting struct {
	gorm.Model
	Key   string `gorm:"size:32;unique_index;not null"`
	Value string `sql:"type:text;"`
}

func GetSettingValues() []Setting {
	var s []Setting
	GetConnection().Model(Setting{}).Scan(&s)
	return s
}

func UpdateSettingValue(s Setting) error {
	c := GetConnection()
	return c.Model(&s).Where(Setting{Key: s.Key}).Assign(Setting{Value: s.Value}).FirstOrCreate(&s).Error
}

