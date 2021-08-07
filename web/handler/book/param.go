package book

import model "github.com/zhenhua32/xingkong/pkg/model/gorm"

type GetBookDirectoryResp struct {
	Total    int                `json:"total"`
	DataList *model.ChapterList `json:"data_list"`
}
