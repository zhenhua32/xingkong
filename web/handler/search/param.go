package search

import (
	s "github.com/zhenhua32/xingkong/pkg/search"
)

type SearchReq struct {
	Keyword string `json:"keyword" binding:"required" validate:"min=1,max=10"`
	Limit   int    `json:"limit" validate:"min=1,max=100"`
}

type SearchResp struct {
	Total    int              `json:"total"`
	DataList []s.SearchResult `json:"data_list"`
}
