package models

import (
	"go-web-scaffold/tools"
	"time"
)

type UserModel struct {
	ID        uint   `gorm:"primary_key"`
	Username  string `gorm:"type:varchar(64);not null"`
	Num       string `gorm:"type:varchar(20);not null"`
	Ua        string `gorm:"type:varchar(256)"`
	Title     string `gorm:"type:varchar(128)"`
	Hash      uint64 `gorm:"unique_index:hash_idx;"`
	CreatedAt time.Time
}

func FindOneUser(condition interface{}) (UserModel, error) {
	db := tools.GetDB()
	var model UserModel
	err := db.Where(condition).First(&model).Error
	return model, err
}

func SaveOne(data interface{}) error {
	db := tools.GetDB()
	err := db.Save(data).Error
	return err
}

func (model *UserModel) Update(data interface{}) error {
	db := tools.GetDB()
	err := db.Model(model).Update(data).Error
	return err
}
