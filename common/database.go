package common

import (
	"go-web-scaffold/models"
	"github.com/jinzhu/gorm"
)

func DBMigrate(db *gorm.DB, flag bool) {
	if flag == true {
		AutoMigrate(db)
	}
}

func DBTableInit(db *gorm.DB, model interface{}) {
	if !db.HasTable(model) {
		if err := db.Set("gorm:table_options",
			"ENGINE=InnoDB DEFAULT CHARSET=utf8").CreateTable(model).Error; err != nil {
			panic(err)
		}
	}
}

func AutoMigrate(db *gorm.DB) {
	// 数据model在这个函数同步
	DBTableInit(db, &models.UserModel{})
	db.AutoMigrate(&models.UserModel{})

}
