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

type ChapterList []Chapter

// UpsertBookChapters 将 book.ChapterList 保存起来, 默认不更新 content 字段
func UpsertBookChapters(b *Book, cl *book.ChapterList, fields []string) (*ChapterList, error) {
	if fields == nil {
		fields = []string{"name", "url", "book_id"}
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

	DB.Omit("content").Where("book_id = ?", b.ID).Find(&chapters)
	return &chapters, err
}
