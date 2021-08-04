package book

import (
	"strings"

	"github.com/gocolly/colly"
	"github.com/gocolly/colly/extensions"

	search_impl "github.com/zhenhua32/xingkong/internal/search"
	"github.com/zhenhua32/xingkong/pkg/book"
	"github.com/zhenhua32/xingkong/pkg/logger"
	"github.com/zhenhua32/xingkong/pkg/search"
)

func NewBook01(s search.SearchResult) *book.Book {
	b := book.Book{
		Name:           s.BookName,
		Author:         s.Author,
		Brief:          s.Brief,
		Url:            s.Url,
		BookType:       s.BookType,
		ImgUrl:         s.ImgUrl,
		LastUpdateTime: s.LastUpdateTime,
		Source:         s.Source,
	}
	b.GetChapterList = GenGetChapterList01(s.Url, &b)

	return &b
}

func NewChapter01(name string, url string, b *book.Book) *book.Chapter {
	c := book.Chapter{
		Name: name,
		Url:  url,
		Book: b,
	}
	c.GetContent = GenGetContent01(url)

	return &c
}

// 返回 GetChapterList 函数
func GenGetChapterList01(url string, b *book.Book) book.GetChapterList {

	return func() (book.ChapterList, error) {
		result := make(book.ChapterList, 0, 100)
		var err error

		// 初始化一个新请求
		c := colly.NewCollector()
		extensions.RandomUserAgent(c)
		extensions.Referer(c)

		c.OnError(func(_ *colly.Response, e error) {
			logger.Sugar.Debug(e)
			err = e
		})

		c.OnHTML(`#list > dl`, func(e *colly.HTMLElement) {
			e.ForEach(`dd`, func(_ int, s *colly.HTMLElement) {
				name := s.ChildText(`a`)
				u, _ := search_impl.BaseUrl01.Parse(s.ChildAttr(`a`, `href`))

				result = append(result, *NewChapter01(name, u.String(), b))
			})
		})

		c.Visit(url)

		return result, err
	}
}

// 返回 GetContent 函数
func GenGetContent01(url string) book.GetContent {
	return func() (string, error) {
		var result string
		var err error

		// 初始化一个新请求
		c := colly.NewCollector()
		extensions.RandomUserAgent(c)
		extensions.Referer(c)

		c.OnError(func(_ *colly.Response, e error) {
			logger.Sugar.Debug(e)
			err = e
		})

		c.OnHTML(`#content`, func(e *colly.HTMLElement) {
			lines := strings.FieldsFunc(e.Text, func(r rune) bool { return r == 160 })
			result = strings.Join(lines, "\n")
		})

		c.Visit(url)

		return result, err
	}
}

func init() {
	book.GlobalBookManagerInstance.RegisterNewBook(search_impl.BaseUrlStr01, NewBook01)
	book.GlobalBookManagerInstance.RegisterNewChapter(search_impl.BaseUrlStr01, NewChapter01)
}
