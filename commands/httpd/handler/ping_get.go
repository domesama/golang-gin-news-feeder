package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//func PingGet(c *gin.Context){
//	c.JSON(http.StatusOK, map[string] string{
//		"owo": "OwO",
//	})
//}

func PingGet() gin.HandlerFunc{ //Higher order allows us to pass dependencies
	return func(c *gin.Context){
		c.JSON(http.StatusOK, map[string]string{
			"OwO": "OwO",
		})
	}
}
