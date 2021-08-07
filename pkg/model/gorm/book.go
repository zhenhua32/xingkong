package model

import (
	"github.com/zhenhua32/xingkong/pkg/book"
	"github.com/zhenhua32/xingkong/pkg/logger"
	"github.com/zhenhua32/xingkong/pkg/search"
	"gorm.io/gorm/clause"
)

type Book struct {
	BaseModel
	book.Book

	// 定义 Has Many 关系
	ChapterList []Chapter `json:"chapter_list"  gorm:"foreignKey:BookID"`
}

type BookList []Book

// UpsertBookSearchResult 根据搜索结果插入或更新数据
func UpsertBookSearchResult(sl *search.SearchResultList) (*BookList, error) {
	var books = make(BookList, 0, 10)
	var urls []string

	for _, s := range *sl {
		b := *book.GBM.NewBook(&s)
		books = append(books, Book{
			Book: b,
		})
		urls = append(urls, b.Url)
	}

	// TODO: 坑爹玩意, upsert 的时候 id 不正确, 参考 https://github.com/go-gorm/gorm/issues/4093
	err := DB.Clauses(clause.OnConflict{
		UpdateAll: true,
	}).Create(&books).Error
	if err != nil {
		logger.Sugar.Errorf("插入小说错误: %v", err)
	}

	// 重新在查一遍吧
	DB.Where("url in ?", urls).Find(&books)
	return &books, err
}
