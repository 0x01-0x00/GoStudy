package gorm

import (
	"GoStudy/gorm/model"
	"fmt"
	"gorm.io/gorm"
)

func DBDemo1(db *gorm.DB) {
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
	// Update - 更新多个字段
	db.Model(&product).Updates(model.Product{Price: 99, Code: "F42"}) // 仅更新非零值字段
	db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})
	fmt.Println("Update End!")

	// Delete
	db.Delete(&model.Product{}, 4)
}
