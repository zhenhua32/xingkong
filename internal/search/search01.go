package search

import (
	"fmt"
	"net/url"
	"time"

	"github.com/gocolly/colly"
	"github.com/gocolly/colly/extensions"
	"github.com/zhenhua32/xingkong/configs"
	"github.com/zhenhua32/xingkong/pkg/search"
)

type SearchEngine01 struct {
	baseUrl string
}

func (g SearchEngine01) String() string {
	return "58小说网"
}

func (g SearchEngine01) Source() string {
	return g.baseUrl
}

func (g SearchEngine01) Search(keyword string, limit int) (search.SearchResultList, error) {
	// 定义返回结果
	result := make(search.SearchResultList, 0, 10)

	// 初始化一个新请求
	c := colly.NewCollector()
	extensions.RandomUserAgent(c)
	extensions.Referer(c)

	// 获取搜索结果
	c.OnHTML("body > div.result-list", func(e *colly.HTMLElement) {
		// 获取每一个搜索结果
		e.ForEach(`.result-item`, func(_ int, e *colly.HTMLElement) {
			baseUrl, _ := url.Parse(g.baseUrl)

			bookName := e.ChildText(`div.result-game-item-detail > h3 > a`)
			author := e.ChildText(`div.result-game-item-detail > div > p:nth-child(1) > span:nth-child(2)`)
			brief := e.ChildText(`div.result-game-item-desc`)

			u, _ := baseUrl.Parse(e.ChildAttr(`div.result-game-item-detail > h3 > a`, "href"))

			bookType := e.ChildText(`div.result-game-item-detail > div > p:nth-child(2) > span:nth-child(2)`)

			imgUrl, _ := baseUrl.Parse(e.ChildAttr(`div.result-game-item-pic > a`, "href"))

			lastUpdateTimeS := e.ChildText(`div.result-game-item-detail > div > p:nth-child(3) > span:nth-child(2)`)
			lastUpdateTime, _ := time.ParseInLocation("2006-01-02 15:04:05", lastUpdateTimeS, configs.TimeZone)

			lastChapter := e.ChildText(`div.result-game-item-detail > div > p:nth-child(4) > a`)

			result = append(result, search.SearchResult{
				BookName:       bookName,
				Author:         author,
				Brief:          brief,
				Url:            u.String(),
				BookType:       bookType,
				ImgUrl:         imgUrl.String(),
				LastUpdateTime: lastUpdateTime,
				LastChapter:    lastChapter,
			})

		})
	})
	// 执行
	c.Visit(fmt.Sprintf("%s/search.php?q=%s", g.baseUrl, keyword))

	return result, nil
}

func init() {
	engine := SearchEngine01{baseUrl: "http://www.wbxsw.com"}
	search.GlobalSearchEngineInstance.Register(&engine)
}
