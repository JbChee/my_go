package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//func main() {
//
//	r := gin.Default()
//	r.LoadHTMLFiles("templates/posts/index.html", "templates/users/index.html")
//	//r.LoadHTMLGlob("templates/**/*")
//
//	r.GET("/posts/index", func(context *gin.Context) {
//		context.HTML(200,"posts/index.html",gin.H{
//			"title":"posts/index",
//		})
//	})
//
//	r.GET("/users/index", func(context *gin.Context) {
//		context.HTML(200,"users/index.html",gin.H{
//			"title":"users/index",
//		})
//	})
//
//	r.Run(":9999")
//
//
//}

func main() {
	r := gin.Default()

	r.Static("/static","./static")
	//r.LoadHTMLGlob("templates/**/*")
	r.LoadHTMLFiles("templates/posts/index.html", "templates/users/index.html")
	r.GET("/posts/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "posts/index.html", gin.H{
			"title": "posts/index",
		})
	})

	r.GET("users/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "users/index.html", gin.H{
			"title": "users/index",
		})
	})

	r.Run(":8080")
}