package model

import (
	"github.com/zhenhua32/xingkong/pkg/book"
	"github.com/zhenhua32/xingkong/pkg/logger"
	"github.com/zhenhua32/xingkong/pkg/search"
	"gorm.io/gorm/clause"
)

type Book struct {
	BaseModel
	Book book.Book `gorm:"embedded" json:"book"`

	// 定义 Has Many 关系
	ChapterList []Chapter `json:"chapter_list"  gorm:"foreignKey:BookID"`
}

type BookList []Book

// UpsertBookSearchResult 根据搜索结果插入或更新数据
func UpsertBookSearchResult(sl *search.SearchResultList) (*BookList, error) {
	var books = make(BookList, 0, 10)

	for _, s := range *sl {
		b := *book.GBM.NewBook(&s)
		books = append(books, Book{
			Book: b,
		})
	}

	err := DB.Clauses(clause.OnConflict{
		UpdateAll: true,
	}).Create(&books).Error
	if err != nil {
		logger.Sugar.Errorf("插入小说错误: %v", err)
	}
	return &books, err
}
