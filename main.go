package main

import (
	"GoStudy/gin"
	"GoStudy/gorm"
	"fmt"
)

var err error

func main() {
	db, _ := gorm.DBInit()
	fmt.Println(db)
	//gorm.DBDemo1(db)

	gin.RouterInit()
}
