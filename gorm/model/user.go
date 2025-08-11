package model

import (
	"database/sql"
	"time"
)

// UserTest 测试结构体
type UserTest struct {
	ID           uint           //主键的标准字段
	Name         string         //常规字符串字段
	Email        *string        //指向字符串的指针，允许空值
	Age          uint8          //一个无符号的8位整数
	Birthday     *time.Time     //指向时间的指针。时间，可以为空
	MemberNumber sql.NullString //使用sql。NullString用于处理可为null的字符串
	ActivatedAt  sql.NullTime   //使用sql。空时间字段的NullTime
	CreatedAt    time.Time      //创建时间由GORM自动管理
	UpdatedAt    time.Time      //由GORM自动管理更新时间
	ignored      string         //未导出的字段将被忽略
}
