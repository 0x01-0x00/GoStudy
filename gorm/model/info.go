package model

import "gorm.io/gorm"

type Info struct {
	gorm.Model
	Address   string
	StudentId uint
}
