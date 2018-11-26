package tools

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Database struct {
	*gorm.DB
}

var DB *gorm.DB

func GetDB() *gorm.DB {
	return DB
}

func DBInit() *gorm.DB {
	//mysqlArgs := "hatlonely:keaiduo1@/hatlonely?charset=utf8&parseTime=True&loc=Local"
	mysqlConnect := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True",
		ConfigStr("mysql.user"),
		ConfigStr("mysql.password"),
		ConfigStr("mysql.host"),
		ConfigStr("mysql.port"),
		ConfigStr("mysql.db"))
	db, err := gorm.Open("mysql", mysqlConnect)
	if err != nil {
		fmt.Println("db err: ", err)
	}
	db.DB().SetMaxIdleConns(100)
	//db.LogMode(true)
	DB = db
	return DB
}
