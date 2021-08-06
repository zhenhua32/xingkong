package model

import (
	"github.com/zhenhua32/xingkong/pkg/book"
)

type Chapter struct {
	BaseModel
	Chapter book.Chapter `gorm:"embedded"`
	Content string       `gorm:"type:text"`
	BookID  uint
}
