package main

import (
	"fmt"
	"os"

	_ "github.com/zhenhua32/xingkong/internal/bookimpl"
	_ "github.com/zhenhua32/xingkong/internal/searchimpl"
	"github.com/zhenhua32/xingkong/pkg/book"
	"github.com/zhenhua32/xingkong/pkg/search"
)

func main() {
	g := search.GSE
	fmt.Println("注册的引擎数量", len(g.EngineList()))

	resultList, _ := g.Search("大奉打更人", 10)

	fmt.Println("结果数量", len(resultList))
	// for _, r := range resultList {
	// 	v := reflect.ValueOf(r)
	// 	for i := 0; i < v.NumField(); i++ {
	// 		fmt.Println(v.Type().Field(i).Name, v.Field(i).Interface())
	// 	}
	// }

	result1 := resultList[0]
	fmt.Println(result1)

	b := book.GBM.NewBook(result1)

	fmt.Println(b)

	clist, e := b.GetChapterList()
	if e != nil {
		fmt.Println(e)
	}

	fmt.Println(len(clist))

	c := clist[0]
	fmt.Println(c)

	content, e := c.GetContent()
	if e != nil {
		fmt.Println(e)
	}
	fmt.Println(content)

	os.WriteFile("a.txt", []byte(content), 0644)
}
