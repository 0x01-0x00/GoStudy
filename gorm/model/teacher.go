package model

import "gorm.io/gorm"

type Teacher struct {
	gorm.Model
	Name    string // 教师姓名
	Number  string // 工号
	ClassId uint   // 班级id
}
