package book

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/zhenhua32/xingkong/pkg/errno"
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
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		handler.JSON(c, errno.NewErrno(errno.ErrBind.Code, c.Param("id")+"不是有效的ID", err), nil)
		return
	}

	b := &model.Book{}
	if err := b.ById(id); err != nil {
		handler.JSON(c, errno.New(errno.ErrNotFound, err), nil)
		return
	}

	handler.JSON(c, nil, b)
}
