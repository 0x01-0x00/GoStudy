package main

import (
	"GoStudy/gorm"
)

func main() {
	db, _ := gorm.DBInit()

	gorm.DBDemo1(db)
}
