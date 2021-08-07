package book

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/zhenhua32/xingkong/pkg/book"
	"github.com/zhenhua32/xingkong/pkg/errno"
	model "github.com/zhenhua32/xingkong/pkg/model/gorm"
	"github.com/zhenhua32/xingkong/web/handler"
)

func getBookFromPath(c *gin.Context) (*model.Book, bool) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		handler.JSON(c, errno.NewErrno(errno.ErrBind.Code, c.Param("id")+"不是有效的ID", err), nil)
		return nil, false
	}

	b := &model.Book{}
	if err := model.DB.First(b, id).Error; err != nil {
		handler.JSON(c, errno.New(errno.ErrNotFound, err), nil)
		return nil, false
	}

	b.Book = *book.GBM.NewBook(&b.Book)

	return b, true
}
