package dao

import "github.com/jinzhu/gorm"

var (
	db *gorm.DB
)

func DB() *gorm.DB {
	return db
}
