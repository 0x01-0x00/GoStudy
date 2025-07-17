package model

import "gorm.io/gorm"

type Student struct {
	gorm.Model
	Name   string // 姓名
	Age    int    // 年龄
	Number string // 学号
	// 默认外键
	ClassID uint // 班级id
	// `Student` belongs to `Class`, `ClassID` is the foreign key
	Class Class

	// `Student` has one `Info`
	Info Info

	// `Student` has many `Hobby`
	Hobbies []Hobby

	// 嵌入方式1
	Score // 成绩
}

type StudentA struct {
	gorm.Model
	Name   string // 姓名
	Age    int    // 年龄
	Number string // 学号
	// 重写外键,重写引用
	ClassRefer uint // 班级id
	// `Student` belongs to `Class`, `ClassRefer` is the foreign key
	Class Class `gorm:"foreignKey:ClassRefer;references:Uid"`

	// 嵌入方式2
	Score Score `gorm:"embedded"` // 成绩
}
