package main

import (
	"gotest1/httpd/handler"
	"gotest1/platform/newsfeed"

	"github.com/gin-gonic/gin"
)

func main() {
	// To insert DB here

	feed := newsfeed.New()

	r := gin.Default()

	r.GET("/newsfeed", handler.NewsFeedGet(feed))
	r.POST("/newsfeed", handler.NewsFeedPost(feed))
	r.Run() // listen and serve on 0.0.0.0:8080
}
