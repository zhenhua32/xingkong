package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Router() http.Handler {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	return r
}
