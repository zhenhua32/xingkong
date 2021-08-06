package search

import (
	"github.com/gin-gonic/gin"
	"github.com/zhenhua32/xingkong/pkg/errno"
	model "github.com/zhenhua32/xingkong/pkg/model/gorm"
	s "github.com/zhenhua32/xingkong/pkg/search"
	"github.com/zhenhua32/xingkong/web/handler"
)

// Search 返回小说的搜索结果
// @Summary 返回小说的搜索结果
// @Description 返回小说的搜索结果
// @ID search
// @Tags 搜索
// @Accept  json
// @Produce  json
// @Param body body search.SearchReq true "搜索参数"
// @Success 200 {object} search.SearchResp
// @Router /search [post]
func Search(c *gin.Context) {
	var req SearchReq
	if err := c.ShouldBindJSON(&req); err != nil {
		handler.JSON(c, errno.NewErrno(errno.ErrBind.Code, err.Error(), err), nil)
		return
	}
	if err := req.Validate(); err != nil {
		handler.JSON(c, errno.NewErrno(errno.ErrBind.Code, err.Error(), err), nil)
		return
	}

	if req.Limit <= 0 {
		req.Limit = 10
	}
	resultList, err := s.GSE.Search(req.Keyword, req.Limit)
	if err != nil {
		handler.JSON(c, err, nil)
	}
	// 将结果保存到数据库中
	model.UpsertBookSearchResult(&resultList)

	handler.JSON(c, nil, SearchResp{Total: len(resultList), DataList: resultList})
}
