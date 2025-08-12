package model

type Student struct {
	ID          uint
	StudentName string // 姓名
	Age         int    // 年龄
	Number      string // 学号
	ClassID     uint
	Class       Class
}
