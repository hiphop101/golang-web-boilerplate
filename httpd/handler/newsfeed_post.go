package handler

import (
	"fmt"
	"golangwebboilerplate/platform/newsfeed"
	"net/http"

	"github.com/gin-gonic/gin"
)

type newsfeedPostRequest struct {
	Title string `json:title`
	Post  string `json:post`
}

func NewsFeedPost(feed newsfeed.Adder) gin.HandlerFunc {
	return func(c *gin.Context) {
		requestBody := newsfeedPostRequest{}

		c.Bind(&requestBody)

		fmt.Println("requestBody=", requestBody)

		item := newsfeed.Item{
			Title: requestBody.Title,
			Post:  requestBody.Post,
		}
		feed.Add(item)

		//c.Status(http.StatusNoContent)
		c.JSON(http.StatusOK, map[string]string{
			"message": "Successfully added",
		})
	}
}
