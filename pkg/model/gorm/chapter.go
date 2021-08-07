package model

import (
	"github.com/zhenhua32/xingkong/pkg/book"
	"github.com/zhenhua32/xingkong/pkg/logger"
	"gorm.io/gorm/clause"
)

type Chapter struct {
	BaseModel
	book.Chapter
	Content string `gorm:"type:text" json:"content"`
	BookID  uint   `json:"book_id"`
}

// Fill 填充, 用于从数据库里取数据后
func (g *Chapter) Fill() error {
	// 填充 Book 字段
	var b = Book{}
	if err := DB.First(&b, g.BookID).Error; err != nil {
		return err
	}

	// 填充 Chapter 字段
	g.Chapter = *book.GBM.NewChapter(g.Name, g.Url, &b.Book, g.Index)
	return nil
}

type ChapterList []Chapter

// TODO: 加一个章节顺序字段
// UpsertBookChapters 将 book.ChapterList 保存起来, 默认不更新 content 字段
func UpsertBookChapters(b *Book, cl *book.ChapterList, fields []string) (*ChapterList, error) {
	if fields == nil {
		fields = []string{"name", "url", "book_id", "index"}
	}

	var chapters = make(ChapterList, 0)

	for _, c := range *cl {
		chapters = append(chapters, Chapter{
			Chapter: c,
			BookID:  b.ID,
		})
	}

	err := DB.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "id"}},
		DoUpdates: clause.AssignmentColumns(fields),
	}).Create(&chapters).Error
	if err != nil {
		logger.Sugar.Errorf("插入章节错误: %v", err)
	}

	DB.Omit("content").Where("book_id = ?", b.ID).Order("index").Find(&chapters)
	return &chapters, err
}
