package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Binding from JSON
type Login struct {
	User     string `form:"user" json:"user" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

func main() {
	r := gin.Default()

	//get 参数
	r.GET("/user/get", func(c *gin.Context) {

		//username := c.DefaultQuery("username","inke")
		username := c.Query("username")
		addr:= c.Query("addr")
		c.JSON(200,gin.H{
			"username":username,
			"addr":addr,
		})
	})

	//form 参数
	r.POST("user/post", func(c *gin.Context) {

		username := c.PostForm("username")
		addr := c.PostForm("address")
		c.JSON(200,gin.H{
			"message": "ok",
			"username":username,
			"addr":addr,

		})
	})

	//提取 url 路径
	r.GET("/user/path/:username/:address", func(c *gin.Context) {
		username := c.Param("username")
		address := c.Param("address")
		//输出json结果给调用方
		c.JSON(http.StatusOK, gin.H{
			"message":  "ok",
			"username": username,
			"address":  address,
		})
	})



	//参数绑定  万能 query  form  json
	r.POST("/loginJSON", func(c *gin.Context) {
		var login Login

		if err := c.ShouldBind(&login); err == nil {
			fmt.Printf("login info:%#v\n", login)
			c.JSON(http.StatusOK, gin.H{
				"user":     login.User,
				"password": login.Password,
			})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
	})



	r.Run(":9999")
}
