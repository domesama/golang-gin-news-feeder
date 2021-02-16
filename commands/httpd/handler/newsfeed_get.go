package handler

import (
	"github.com/domesama/golang-gin-news-feeder/packages/newsfeed"
	"github.com/gin-gonic/gin"
	"net/http"
)

//func NewsfeedGet(feed *newsfeed.Repo) gin.HandlerFunc{
//	return func(c *gin.Context) {
//		c.JSON(http.StatusOK, feed.GetAll())
//	}
//}

func NewsfeedGet(feed newsfeed.Getter) gin.HandlerFunc{
	return func(c *gin.Context){
		c.JSON(http.StatusOK, feed.GetAll())
	}
}