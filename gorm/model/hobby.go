package model

import "gorm.io/gorm"

type Hobby struct {
	gorm.Model
	Read       bool
	Write      bool
	Football   bool
	Basketball bool
	StudentId  uint
}
