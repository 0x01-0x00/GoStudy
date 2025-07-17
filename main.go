package main

import (
	"GoStudy/gorm"
	"fmt"
)

func main() {
	db, _ := gorm.DBInit()
	fmt.Println(db)
	//gorm.DBDemo1(db)
}
