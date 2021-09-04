package xdb

import (
	"net/http"
	"sy/models"
	"sy/pkg/e"

	"github.com/gin-gonic/gin"
)

func Encode(c *gin.Context) {
	var config Config
	code := e.INVALID_PARAMS
	if c.BindJSON(&config) == nil {
		code = e.SUCCESS
		//将数据加入到异步队列
		if content, ok := config.str(); ok {
			models.Push("1", content)
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"msg":    e.GetMsg(code),
	})

	c.Abort()
}

func Decode(c *gin.Context) {
	var config Config
	code := e.INVALID_PARAMS
	if c.BindJSON(&config) == nil {
		code = e.SUCCESS
		//将数据加入到异步队列
		if content, ok := config.str(); ok {
			models.Push("6", content)
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"status": code,
		"msg":    e.GetMsg(code),
	})

	c.Abort()
}
