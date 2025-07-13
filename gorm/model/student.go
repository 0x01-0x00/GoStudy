package model

import "gorm.io/gorm"

type Student struct {
	gorm.Model
	Name    string // 姓名
	Age     int    // 年龄
	Number  string // 学号
	ClassId uint   // 班级id
	// 嵌入方式1
	//Score          // 成绩
	// 嵌入方式2
	Score Score `gorm:"embedded"` // 成绩
}
