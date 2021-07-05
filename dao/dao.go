package dao

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var (
	db *gorm.DB
)

func DB() *gorm.DB {
	return db
}

func MustInitDB(username string, password string, hostname string, hostport string, database string) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", username, password, hostname, hostport, database)
	d, err := gorm.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	err = d.DB().Ping()
	if err != nil {
		panic(err)
	}
	db = d
}
