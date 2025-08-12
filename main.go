package main

import (
	"GoStudy/gorm"
	"fmt"
)

var err error

func main() {
	db, _ := gorm.DBInit()
	fmt.Println("db: ", db)
	//gorm.DBDemo1(db, err)
	gorm.DBDemo2(db, err)
}
