package main

import (
	"GoStudy/gorm"
	"GoStudy/gorm/model"
	"fmt"
)

func main() {
	db, _ := gorm.DBInit()

	//// Create
	//// 创建一条记录
	//db.Create(&model.Product{Code: "D42", Price: 100})
	fmt.Println("Create End!")

	// Read
	// 检索单个对象
	var product model.Product
	db.First(&product, 1)                 // 根据整型主键查找
	db.First(&product, "code = ?", "D42") // 查找 code 字段值为 D42 的记录
	fmt.Println(product)
	fmt.Println("Read End!")

	// 检索全部对象
	var products []model.Product
	db.Find(&products, "price = ?", 100) // 查找 price 字段值为 100 的记录
	fmt.Println(products)
	fmt.Println("Find End!")

	// Update
	db.Model(&product).Update("Price", 99) // 当使用 Model 方法，并且它有主键值时，主键将会被用于构建条件
	fmt.Println("Update End!")
}
