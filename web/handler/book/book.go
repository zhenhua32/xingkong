package book

import (
	"github.com/gin-gonic/gin"
	model "github.com/zhenhua32/xingkong/pkg/model/gorm"
	"github.com/zhenhua32/xingkong/web/handler"
)

// GetBook 返回小说的详情
// @Summary 返回小说的详情
// @Description 返回小说的详情
// @Tags book
// @Accept  text/html
// @Produce  json
// @Param id path uint64 true "小说ID"
// @Success 200 {object} model.Book
// @Router /book/{id} [get]
func GetBook(c *gin.Context) {
	b, ok := getBookFromPath(c)
	if !ok {
		return
	}

	handler.JSON(c, nil, b)
}

// GetBookDirectory 返回小说的目录
// @Summary 返回小说的目录
// @Description 返回小说的目录
// @Tags book
// @Accept  text/html
// @Produce  json
// @Param id path uint64 true "小说ID"
// @Success 200 {object} book.GetBookDirectoryResp
// @Router /book/{id}/directory [get]
func GetBookDirectory(c *gin.Context) {
	b, ok := getBookFromPath(c)
	if !ok {
		return
	}

	// 它从数据库里恢复, 还没有这个 GetChapterList 方法
	chapterList, err := b.Book.GetChapterList()
	if err != nil {
		handler.JSON(c, err, nil)
		return
	}

	// 保存章节结果
	cl, err := model.UpsertBookChapters(b, &chapterList, nil)
	if err != nil {
		handler.JSON(c, err, nil)
		return
	}

	handler.JSON(c, nil, GetBookDirectoryResp{Total: len(*cl), DataList: cl})
}
