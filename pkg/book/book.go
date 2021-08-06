/*
	book 包定义了基础的书籍模型
*/
package book

import (
	"time"

	"github.com/zhenhua32/xingkong/pkg/logger"
	"github.com/zhenhua32/xingkong/pkg/search"
)

// Book 定义了书籍信息
type Book struct {
	Name           string    `json:"name" gorm:"type:string;size:64"`             // 书名
	Author         string    `json:"author" gorm:"type:string;size:64"`           // 作者
	Brief          string    `json:"brief" gorm:"type:text"`                      // 简介
	Url            string    `json:"url" gorm:"type:string;size:256;uniqueIndex"` // 链接
	BookType       string    `json:"book_type" gorm:"type:string;size:16"`        // 类型
	ImgUrl         string    `json:"img_url" gorm:"type:string;size:256"`         // 图片链接
	LastUpdateTime time.Time `json:"last_update_time"`                            // 最近更新时间
	LastChapter    *Chapter  `json:"last_chapter" gorm:"-"`                       // 最近更新章节
	Source         string    `json:"source" gorm:"type:string;size:256"`          // 来源

	// 定义对应的方法
	GetChapterList GetChapterList `json:"-" gorm:"-"`
}

// Chapter 定义了章节信息
type Chapter struct {
	Name string `json:"name" gorm:"type:string;size:256"`            // 章节名称
	Url  string `json:"url" gorm:"type:string;size:256;uniqueIndex"` // 链接
	Book *Book  `json:"-" gorm:"-"`                                  // 书籍

	// 定义对应的方法
	GetContent GetContent `json:"-" gorm:"-"`
}

type ChapterList []Chapter

// GetChapterList 定义获取所有章节列表的函数
type GetChapterList func() (ChapterList, error)

// GetContent 定义获取章节内容的函数
type GetContent func() (string, error)

type NewBook func(s search.SearchResult) *Book
type NewChapter func(name string, url string, book *Book) *Chapter

// 全局书籍管理器
type GlobalBookManager struct {
	bookFuncMap   map[string]NewBook
	chaperFuncMap map[string]NewChapter
}

// 在源上注册 NewBook 函数
func (g *GlobalBookManager) RegisterNewBook(source string, f NewBook) {
	g.bookFuncMap[source] = f
}

// 在源上注册 NewChapter 函数
func (g *GlobalBookManager) RegisterNewChapter(source string, f NewChapter) {
	g.chaperFuncMap[source] = f
}

// NewBook 将 SearchResult 转换成 Book
func (g *GlobalBookManager) NewBook(s search.SearchResult) *Book {
	f := g.bookFuncMap[s.Source]
	if f == nil {
		logger.Sugar.Infof("找不到对应的 NewBook 函数, 源是 %s", s.Source)
		return nil
	}

	return f(s)
}

// NewChapter 获取一个 Chapter
func (g *GlobalBookManager) NewChapter(name string, url string, book *Book) *Chapter {
	f := g.chaperFuncMap[book.Source]
	if f == nil {
		logger.Sugar.Infof("找不到对应的 NewChapter 函数, 源是 %s", book.Source)
		return nil
	}

	return f(name, url, book)
}

var GBM = &GlobalBookManager{
	bookFuncMap:   make(map[string]NewBook),
	chaperFuncMap: make(map[string]NewChapter),
}
