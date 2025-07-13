package main

import (
	"GoStudy/gorm"
	"GoStudy/gorm/model"
)

func main() {
	db, _ := gorm.DBInit()

	// Create
	db.Create(&model.Product{Code: "D42", Price: 100})

}
