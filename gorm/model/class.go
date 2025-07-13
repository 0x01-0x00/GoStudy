package model

import "gorm.io/gorm"

type Class struct {
	gorm.Model
	Grade  string // 年级
	Number uint   // 班级编号
}
