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

	// base
	r.GET("/", apis.Root)
	r.GET("/ping", apis.Ping)
	r.GET("/author", apis.Author)

	// index
	r.GET("/index.html", apis.Index)

	return r
}
