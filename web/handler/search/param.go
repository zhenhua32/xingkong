package search

import (
	model "github.com/zhenhua32/xingkong/pkg/model/gorm"
	"github.com/zhenhua32/xingkong/web/valid"
)

type SearchReq struct {
	Keyword string `json:"keyword" validate:"required,min=1,max=20" minLength:"1" maxLength:"20"` // 搜索关键字
	Limit   int    `json:"limit" validate:"min=1,max=100" minimum:"1" maximum:"100" default:"10"` // 限制结果数量
}

func (g *SearchReq) Validate() error {
	return valid.Validate.Struct(g)
}

type SearchResp struct {
	Total    int             `json:"total"`
	DataList *model.BookList `json:"data_list"`
}
