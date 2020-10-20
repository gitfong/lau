package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/gitfong/lau/apis"
)

func initRouter() *gin.Engine {
	// gin.SetMode(gin.ReleaseMode)
	// gin.DisableConsoleColor()

	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.StaticFS("/statics", http.Dir("./statics"))
	r.StaticFile("/favicon.ico", "./resources/favicon.ico")

	r.GET("/", apis.Index)
	r.GET("/index.html", apis.Index)

	r.GET("/author", apis.Author)

	return r
}
