package handler

import (
	"github.com/domesama/golang-gin-news-feeder/packages/newsfeed"
	"github.com/gin-gonic/gin"
	"net/http"
)

type newsfeedPostRequest struct{
	Title string `json:"title"`
	Post string `json:post`
}

//func NewsfeedPost(feed *newsfeed.Repo) gin.HandlerFunc{
//	return func(c *gin.Context){
//		requestBody := newsfeedPostRequest{}
//		c.Bind(&requestBody)
//
//		item := newsfeed.Item{
//			Title: requestBody.Title,
//			Post: requestBody.Post,
//		}
//		feed.Add(item)
//		c.Status(http.StatusOK)
//	}
//}

func NewsfeedPost(feed newsfeed.Setter) gin.HandlerFunc{
	return func(c *gin.Context){
		requestBody := &newsfeedPostRequest{}
		c.Bind(requestBody)

		item := newsfeed.Item{
			Title: requestBody.Title,
			Post:  requestBody.Post,
		}
		feed.Add(item)

		c.Status(http.StatusAccepted)
	}
}