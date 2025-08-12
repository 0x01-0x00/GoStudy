package gorm

import (
	"database/sql"
	"fmt"
	"time"

	"gorm.io/gorm"
)

// UserTest 测试结构体
type UserTest struct {
	ID           uint           // 主键的标准字段
	Name         string         // 常规字符串字段
	Email        *string        // 指向字符串的指针，允许空值
	Age          uint8          // 一个无符号的8位整数
	Birthday     *time.Time     // 指向时间的指针。时间，可以为空
	MemberNumber sql.NullString // 使用sql。NullString用于处理可为null的字符串
	ActivatedAt  sql.NullTime   // 使用sql。空时间字段的NullTime
	CreatedAt    time.Time      // 创建时间由GORM自动管理
	UpdatedAt    time.Time      // 由GORM自动管理更新时间
	ignored      string         // 未导出的字段将被忽略
}

// BeforeCreate 创建前
func (u *UserTest) BeforeCreate(tx *gorm.DB) (err error) {
	u.Name = u.Name + "_BeforeCreate"
	return
}

func DBDemo1(db *gorm.DB, err error) {
	email := "zhongtao@gmail.com"
	birthday := time.Now()
	user := UserTest{
		ID:           0,
		Name:         "钟涛",
		Email:        &email,
		Age:          27,
		Birthday:     &birthday,
		MemberNumber: sql.NullString{},
		ActivatedAt:  sql.NullTime{},
		CreatedAt:    time.Time{},
		UpdatedAt:    time.Time{},
	}

	//tx := db.Create(&user)
	//fmt.Println(tx.Error)
	//fmt.Println(tx.RowsAffected)

	//users := []UserTest{
	//	{Name: "Jinzhu", Age: 18, Birthday: &birthday},
	//	{Name: "Jackson", Age: 19, Birthday: &birthday},
	//}
	//tx = db.Create(&users)
	//fmt.Println(tx.Error)
	//fmt.Println(tx.RowsAffected)

	db.Debug().First(&user, 0)
	fmt.Println(user)
}
