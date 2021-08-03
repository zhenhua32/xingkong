package main

import (
	"fmt"

	_ "github.com/zhenhua32/xingkong/internal/search"
	"github.com/zhenhua32/xingkong/pkg/search"
)

func main() {
	g := search.GlobalSearchEngineInstance
	fmt.Println("注册的引擎数量", len(g.EngineList()))

	result, _ := g.Search("大奉打更人", 10)

	fmt.Println("结果数量", len(result))
	for _, v := range result {
		fmt.Println(v)
	}
}
