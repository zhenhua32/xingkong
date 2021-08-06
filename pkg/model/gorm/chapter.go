package model

import (
	"github.com/zhenhua32/xingkong/pkg/book"
	"gorm.io/gorm"
)

type Chapter struct {
	gorm.Model
	Chapter book.Chapter `gorm:"embedded"`
	Content string       `gorm:"type:text"`
	BookID  uint
}
