/*
	book 包定义了基础的书籍模型
*/
package book

import "time"

// Book 定义了书籍信息
type Book struct {
	Name           string    `json:"name"`             // 书名
	Author         string    `json:"author"`           // 作者
	Brief          string    `json:"brief"`            // 简介
	Url            string    `json:"url"`              // 链接
	BookType       string    `json:"book_type"`        // 类型
	ImgUrl         string    `json:"img_url"`          // 图片链接
	LastUpdateTime time.Time `json:"last_update_time"` // 最近更新时间
	LastChapter    Chapter   `json:"last_chapter"`     // 最近更新章节

	// 定义对应的方法
	GetChapterList GetChapterList `json:"-"`
}

// Chapter 定义了章节信息
type Chapter struct {
	Name string `json:"name"` // 章节名称
	Url  string `json:"url"`  // 链接
	Book *Book  `json:"-"`    // 书籍

	// 定义对应的方法
	GetContent GetContent `json:"-"`
}

type ChapterList []Chapter
type GetChapterList func() (ChapterList, error)
type GetContent func() (string, error)

type BookManager interface {
	GetChapterList(book *Book) (ChapterList, error)
	GetContent(chapter *Chapter) (string, error)
}
