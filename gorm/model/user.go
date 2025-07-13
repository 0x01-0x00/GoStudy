package model

import (
	"database/sql"
	"time"
)

type User struct {
	ID           uint           // 主键的标准字段
	Name         string         // 主键的标准字段
	Email        *string        // 一个指向字符串的指针，允许为空值
	Age          uint8          // 一个无符号的8位整数
	Birthday     *time.Time     // 一个指向时间的指针。时间，可以为空
	MemberNumber sql.NullString // 使用sql.NullString来处理可为空的字符串
	ActivatedAt  sql.NullTime   // 对于可为空的时间字段，使用sql.NullTime
	CreatedAt    time.Time      // 由GORM自动管理创建时间
	UpdatedAt    time.Time      // 由GORM自动管理更新时间
	ignored      string         // 未导出的字段将被忽略
}
