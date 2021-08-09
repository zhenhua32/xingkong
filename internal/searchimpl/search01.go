package searchimpl

import (
	"fmt"
	"net/url"
	"time"

	"github.com/gocolly/colly"
	"github.com/gocolly/colly/extensions"
	"github.com/zhenhua32/xingkong/configs"
	"github.com/zhenhua32/xingkong/pkg/logger"
	"github.com/zhenhua32/xingkong/pkg/search"
)

var BaseUrlStr01 string = "http://www.wbxsw.com"
var BaseUrl01 *url.URL

type SearchEngine01 struct {
}

func (g SearchEngine01) String() string {
	return "58小说网"
}

func (g SearchEngine01) Source() string {
	return BaseUrlStr01
}

func (g SearchEngine01) Search(keyword string, limit int) (search.SearchResultList, error) {
	// 定义返回结果
	result := make(search.SearchResultList, 0, 10)
	var err error

	// 初始化一个新请求
	c := colly.NewCollector()
	extensions.RandomUserAgent(c)
	extensions.Referer(c)

	c.OnError(func(_ *colly.Response, e error) {
		logger.Sugar.Debug(e)
		err = e
	})

	// 获取搜索结果
	c.OnHTML("body > div.result-list", func(e *colly.HTMLElement) {
		// 获取每一个搜索结果
		e.ForEach(`.result-item`, func(_ int, e *colly.HTMLElement) {

			bookName := e.ChildText(`div.result-game-item-detail > h3 > a`)
			author := e.ChildText(`div.result-game-item-detail > div > p:nth-child(1) > span:nth-child(2)`)
			brief := e.ChildText(`p.result-game-item-desc`)

			u, _ := BaseUrl01.Parse(e.ChildAttr(`div.result-game-item-detail > h3 > a`, "href"))

			bookType := e.ChildText(`div.result-game-item-detail > div > p:nth-child(2) > span:nth-child(2)`)

			imgUrl, _ := BaseUrl01.Parse(e.ChildAttr(`div.result-game-item-pic > a > img`, "src"))

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
				Source:         g.Source(),
			})

		})
	})
	// 执行
	c.Visit(fmt.Sprintf("%s/search.php?q=%s", BaseUrlStr01, keyword))

	return result, err
}

func init() {
	BaseUrl01, _ = url.Parse(BaseUrlStr01)
	engine := SearchEngine01{}
	search.GSE.Register(&engine)
}
