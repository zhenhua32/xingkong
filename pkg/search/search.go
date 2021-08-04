/*
	search 包定义了基础的搜索功能

	SearchEngine 接口公开了特定源应该实现的方法, 然后在 GlobalSearchEngineInstance 上注册
	GlobalSearchEngine 接口定义了全局的搜索引擎, 可以通过 Register 方法聚合各种源
	GlobalSearchEngineInstance 是 GlobalSearchEngine 的唯一实例
*/
package search

import (
	"time"

	"github.com/zhenhua32/xingkong/pkg/logger"
)

// SearchResult 定义了单条搜索结果的信息
type SearchResult struct {
	BookName       string    `json:"book_name"`        // 书名
	Author         string    `json:"author"`           // 作者
	Brief          string    `json:"brief"`            // 简介
	Url            string    `json:"url"`              // 链接
	BookType       string    `json:"book_type"`        // 类型
	ImgUrl         string    `json:"img_url"`          // 图片链接
	LastUpdateTime time.Time `json:"last_update_time"` // 最近更新时间
	LastChapter    string    `json:"last_chapter"`     // 最近更新章节
	Source         string    `json:"source"`           // 来源
}

type SearchResultList []SearchResult

// SearchEngine 在某个特定源上实现了搜索方法
type SearchEngine interface {
	// Search 搜索调用当前的引擎, 并返回搜索结果
	Search(keyword string, limit int) (SearchResultList, error)
	// String 以字符串形式返回当前引擎的信息
	String() string
	// Source 定义当前引擎的数据来源, 也是标识符
	Source() string
}

// 全局搜索引擎
type GlobalSearchEngine struct {
	engineList []SearchEngine
}

func (g *GlobalSearchEngine) EngineList() []SearchEngine {
	return g.engineList
}

// 注册一个引擎
func (g *GlobalSearchEngine) Register(engine SearchEngine) {
	g.engineList = append(g.engineList, engine)
}

// 搜索全局引擎, 会调用所有的引擎的 Search 方法, 并忽略错误
// limit 是每个搜索引擎的搜索数量, 不是总的返回结果数量
func (g *GlobalSearchEngine) Search(keyword string, limit int) (SearchResultList, error) {
	if limit <= 0 {
		limit = 10
	}

	result := make(SearchResultList, 0)
	for _, e := range g.engineList {
		r, err := e.Search(keyword, limit)
		if err != nil {
			// 记录错误
			logger.Sugar.Errorf("search error: %s", err)
			continue
		}
		result = append(result, r...)
	}
	if len(result) > limit {
		result = result[:limit]
	}
	return result, nil
}

var GSE = &GlobalSearchEngine{}
