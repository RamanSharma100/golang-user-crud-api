package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	DBConn *gorm.DB
)

func InitDB() {
	db, err := gorm.Open("mysql", "root:@tcp(localhost:3306)/crudapi?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	DBConn = db

}

func GetDB() *gorm.DB {
	return DBConn
}
