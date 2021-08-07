package chapter

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/zhenhua32/xingkong/pkg/errno"
	model "github.com/zhenhua32/xingkong/pkg/model/gorm"
	"github.com/zhenhua32/xingkong/web/handler"
)

func getChapterFromPath(c *gin.Context) (*model.Chapter, bool) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		handler.JSON(c, errno.NewErrno(errno.ErrBind.Code, c.Param("id")+"不是有效的ID", err), nil)
		return nil, false
	}

	ch := &model.Chapter{}
	if err := model.DB.First(ch, id).Error; err != nil {
		handler.JSON(c, errno.New(errno.ErrNotFound, err), nil)
		return nil, false
	}

	if err := ch.Fill(); err != nil {
		handler.JSON(c, errno.NewErrno(errno.ErrBind.Code, "无法填充章节, 可能是 book 不存在", err), nil)
		return nil, false
	}

	return ch, true
}
