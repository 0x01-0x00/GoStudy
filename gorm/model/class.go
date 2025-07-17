package model

import "gorm.io/gorm"

type Class struct {
	gorm.Model
	Uid    string `gorm:"unique"` // Uid
	Grade  string // 年级
	Number uint   // 班级编号
}
