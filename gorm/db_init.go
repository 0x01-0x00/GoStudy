package gorm

import (
	"GoStudy/gorm/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func DBInit() (db *gorm.DB, err error) {
	// 连接数据库
	dsn := "root:rootroot@tcp(47.94.102.125:3306)/gorm?charset=utf8mb4&parseTime=True&loc=Local"
	// 创建连接
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// 迁移数据结构
	err = db.AutoMigrate(
		&model.UserTest{},
	)
	if err != nil {
		panic("failed to db.AutoMigrate")
	}
	return db, err
}
