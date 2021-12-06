package main

import (
	"github.com/gin-gonic/gin"
)

func f1(c *gin.Context ){
	c.JSON(200,gin.H{
		"message":"hello world",
	})
}


func main() {
	//初始化gin
	r := gin.Default()

	r.GET("/hello",f1)

	r.Run("127.0.0.1:9999")
}
