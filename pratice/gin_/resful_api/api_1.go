package main

import (

	"github.com/gin-gonic/gin"
)


type msg struct{
	Name string `json:name`
	Message string `json:message`
	Age int8 `json:age`

}


func main() {

	r := gin.Default()

	//拼接结构体
	r.GET("/book", func(context *gin.Context) {
		context.JSON(200,gin.H{
			"message":"/book_get",
		})
	})

	r.POST("/book", func(context *gin.Context) {
		context.JSON(200,gin.H{
			"message":"book_post",
		})
	})


	//结构体
	r.GET("/morejson", func(c *gin.Context) {
		var msg msg
		msg.Name = "inke.com"
		msg.Message = "this is a message"
		msg.Age = 18
		c.JSON(200,msg)
	})

	//xml 渲染
	r.GET("/xml", func(c *gin.Context) {

		type msgxml struct{
			Name string `xml:name`
			Message string `xml:message`
			Age int8 `xml:age`

		}
		var msg msgxml
		msg.Name = "inke.com"
		msg.Message = "this is a message"
		msg.Age = 18
		c.XML(200,msg)
	})

	//yaml 渲染
	r.GET("/yaml", func(c *gin.Context) {

		type msgyaml struct{
			Name string `xml:name`
			Message string `xml:message`
			Age int8 `xml:age`

		}
		var msg msgyaml
		msg.Name = "inke.com"
		msg.Message = "this is a message"
		msg.Age = 18
		c.YAML(200,msg)
	})



	r.Run(":9999")





}