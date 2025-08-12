package gorm

import (
	"GoStudy/gorm/model"
	"fmt"

	"gorm.io/gorm"
)

func DBDemo2(db *gorm.DB, err error) {
	//students := []model.Student{
	//	{
	//		StudentName: "大学生1",
	//		Age:         18,
	//		Number:      "202500001",
	//		Class: model.Class{
	//			ClassLevel: "大一",
	//			ClassNum:   "1",
	//			Refer:      "2025001",
	//		},
	//	},
	//	{
	//		StudentName: "大学生2",
	//		Age:         18,
	//		Number:      "202500002",
	//		Class: model.Class{
	//			ClassLevel: "大一",
	//			ClassNum:   "1",
	//			Refer:      "2025001",
	//		},
	//	},
	//}

	//db.Create(&students)

	student := model.Student{}
	// 多对一
	//db.First(&student, 1)
	//fmt.Println(student) // {1 大学生1 18 202500001 1 {0   }}
	//db.Preload("Class").First(&student, 1)
	//fmt.Println(student) // {1 大学生1 18 202500001 1 {1 2025001 大一 1}}

	// 一对一
	db.Preload("StudentInfo").First(&student, 1)
	fmt.Println(student)

	// 一对多
	db.Preload("StudentInfos").First(&student, 1)
	fmt.Println(student)
}
