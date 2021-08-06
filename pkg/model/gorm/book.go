package model

import (
	"github.com/jinzhu/gorm"
	"github.com/zhenhua32/xingkong/pkg/book"
)

type Book struct {
	gorm.Model
	Book book.Book `gorm:"embedded"`

	// 定义 Has Many 关系
	ChapterList []Chapter `json:"chapter_list"  gorm:"foreignKey:BookID"`
}
