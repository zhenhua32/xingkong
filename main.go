package main

import (
	"fmt"
	"math/rand"
	"os"
	"reflect"
	"time"

	"github.com/Rican7/retry"
	"github.com/Rican7/retry/backoff"
	"github.com/Rican7/retry/jitter"
	"github.com/Rican7/retry/strategy"
	_ "github.com/zhenhua32/xingkong/internal/bookimpl"
	_ "github.com/zhenhua32/xingkong/internal/searchimpl"
	"github.com/zhenhua32/xingkong/pkg/book"
	"github.com/zhenhua32/xingkong/pkg/search"
)

func download(r search.SearchResult) {
	b := book.GBM.NewBook(r)

	file, _ := os.OpenFile(r.BookName+".txt", os.O_CREATE|os.O_WRONLY, 0644)

	cList, e := b.GetChapterList()
	if e != nil {
		fmt.Println(e)
	}

	fmt.Println("总章节数:", len(cList))
	fmt.Println("开始下载每一个章节")

	for i, c := range cList {
		fmt.Println(i, c.Name)
		// 写入标题
		file.WriteString(c.Name + "\n")

		var content string

		seed := time.Now().UnixNano()
		random := rand.New(rand.NewSource(seed))
		err := retry.Retry(
			func(attempt uint) error {
				content, e = c.GetContent()
				return e
			},
			strategy.Limit(3),
			strategy.BackoffWithJitter(
				backoff.BinaryExponential(400*time.Millisecond),
				jitter.Deviation(random, 0.5),
			),
		)
		if err != nil {
			fmt.Println("获取失败, 错误是", err)
			continue
		}

		// 写入内容
		file.WriteString(content + "\n\n")
		time.Sleep(time.Second)
	}
}

func printStruct(s interface{}) {
	v := reflect.ValueOf(s)
	for i := 0; i < v.NumField(); i++ {
		fmt.Println(v.Type().Field(i).Name, v.Field(i).Interface())
	}
}

func main() {
	g := search.GSE
	fmt.Println("注册的引擎数量", len(g.EngineList()))

	resultList, _ := g.Search("大奉打更人", 10)

	fmt.Println("结果数量", len(resultList))

	result1 := resultList[0]
	fmt.Println(result1)
	printStruct(result1)

	download(result1)
}
