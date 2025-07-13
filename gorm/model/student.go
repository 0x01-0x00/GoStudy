package model

import "gorm.io/gorm"

type Student struct {
	gorm.Model
	Name    string // 姓名
	Age     int    // 年龄
	Number  string // 学号
	ClassId uint   // 班级id
	Score   Score  // 成绩
}
