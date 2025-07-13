package main

import (
	"GoStudy/gorm"
	"GoStudy/gorm/model"
	"fmt"
)

func main() {
	db, _ := gorm.DBInit()

	//// Create
	//db.Create(&model.Product{Code: "D42", Price: 100})

	// Read
	var product model.Product
	db.First(&product, 1)                 // 根据整型主键查找
	db.First(&product, "code = ?", "D42") // 查找 code 字段值为 D42 的记录
	fmt.Println(product)
}
