package chapter

import (
	"github.com/gin-gonic/gin"
	model "github.com/zhenhua32/xingkong/pkg/model/gorm"
	"github.com/zhenhua32/xingkong/web/handler"
)

// GetChapter 返回章节的详情
// @Summary 返回章节的详情
// @Description 返回章节的详情
// @Tags chapter
// @Accept  text/html
// @Produce  json
// @Param id path uint64 true "章节ID"
// @Success 200 {object} model.Chapter
// @Router /chapter/{id} [get]
func GetChapter(c *gin.Context) {
	chapter, ok := getChapterFromPath(c)
	if !ok {
		return
	}

	// 如果 Content 不存在, 就要保存起来
	if chapter.Content == "" {
		var err error
		chapter.Content, err = chapter.Chapter.GetContent()
		if err != nil {
			handler.JSON(c, err, nil)
			return
		}
		model.DB.Save(chapter)
	}

	handler.JSON(c, nil, chapter)
}
