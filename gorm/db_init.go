package gorm

import (
	"GoStudy/gorm/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func DBInit() (db *gorm.DB, err error) {
	dsn := "root:rootroot@tcp(47.122.87.56:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// 迁移 schema
	err = db.AutoMigrate(
		model.Product{},
	)
	if err != nil {
		panic("failed to db.AutoMigrate")
	}
	return db, err
}
