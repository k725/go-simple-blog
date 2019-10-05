package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/k725/go-simple-blog/config"
)

type DbTarget struct {
	Address  string
	User     string
	Password string
	Database string
}

var dbConn *gorm.DB

// GetConnection ...
func GetConnection() *gorm.DB {
	if dbConn == nil {
		panic("Not initialized database connection.")
	}
	return dbConn
}

// SetupConnection ...
// NOTE: https://github.com/go-sql-driver/mysql
func SetupConnection(c DbTarget) (*gorm.DB, error) {
	con := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=UTC",
		c.User, c.Password, c.Address, c.Database)
	d, err := gorm.Open("mysql", con)
	if err != nil {
		return nil, err
	}
	d.LogMode(config.IsDevelopMode())
	dbConn = d
	return d, nil
}

func CloseConnection() error {
	return dbConn.Close()
}
