package main

import (
	"fmt"
	"github.com/domesama/golang-gin-news-feeder/commands/httpd/handler"
	"github.com/gin-gonic/gin"
)

func main(){
	fmt.Println("Welcome to Golang GIN Framework")

	router := gin.Default() //Basically gin.new() ,but adds the recovery and loggers middlewares

	router.GET("/ping", handler.PingGet())

	router.Run(":8080")
}
