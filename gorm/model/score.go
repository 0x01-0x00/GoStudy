package model

type Score struct {
	//gorm.Model // 被嵌入的结构体不能添加gorm.Model
	Chinese   uint // 语文成绩，范围0~150
	Math      uint // 数学成绩，范围0~150
	English   uint // 英语成绩，范围0~150
	Physics   uint // 物理成绩，范围0~100
	Chemistry uint // 化学成绩，范围0~100
	Biology   uint // 生物成绩，范围0~100
}
