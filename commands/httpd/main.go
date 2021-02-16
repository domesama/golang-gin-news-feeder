package main

import (
	"github.com/domesama/golang-gin-news-feeder/commands/httpd/handler"
	"github.com/domesama/golang-gin-news-feeder/packages/newsfeed"
	"github.com/gin-gonic/gin"
)

func main(){

	feed := newsfeed.New()

	router := gin.Default() //Basically gin.new() ,but adds the recovery and loggers middlewares

	router.GET("/ping", handler.PingGet())
	router.GET("/newsfeed", handler.NewsfeedGet(feed))
	router.POST("/newsfeed", handler.NewsfeedPost(feed))

	router.Run(":8080")

}
