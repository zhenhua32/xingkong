package app

import (
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	_ "github.com/zhenhua32/xingkong/api"
	"github.com/zhenhua32/xingkong/web/handler/book"
	"github.com/zhenhua32/xingkong/web/handler/chapter"
	"github.com/zhenhua32/xingkong/web/handler/ping"
	"github.com/zhenhua32/xingkong/web/handler/search"
)

func loadSwagger(g *gin.Engine) {
	// url := ginSwagger.URL("/swagger/doc.json")
	// g.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
	g.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

func loadAPI(g *gin.Engine) {
	loadSwagger(g)

	g.GET("/ping", ping.Ping)

	g.POST("/search", search.Search)

	g.GET("/book/:id", book.GetBook)
	g.GET("/book/:id/directory", book.GetBookDirectory)

	g.GET("/chapter/:id", chapter.GetChapter)
}
