package model

type Student struct {
	ID          uint
	StudentName string // 姓名
	Age         int    // 年龄
	Number      string // 学号
	// 多对一
	//ClassID     uint   // 默认外键
	//Class       Class
	ClassRefer uint  // 重写外键
	Class      Class `gorm:"foreignKey:ClassRefer"`
	// 一对一
	StudentInfo StudentInfo
	// 一对多
	StudentInfos []StudentInfo
	// 多对多
	Scores []Score `gorm:"many2many:student_score"`
}
