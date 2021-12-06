package routers

import (
	"bolg/global"
	mid "bolg/internal/middleware"
	"bolg/internal/routers/api"
	v1 "bolg/internal/routers/api/v1"
	"bolg/pkg/limiter"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

var methodLimiters = limiter.NewMethodLimiter().AddBuckets(
	limiter.LimiterBucketRule{
		Key:          "/auth",
		FillInterval: time.Second,
		Capacity:     10,
		Quantum:      10,
	},
)

func NewRouter() *gin.Engine {
	r := gin.New()
	if global.ServerSetting.RunMode == "debug" {
		r.Use(gin.Logger())
		r.Use(gin.Recovery())
	} else {
		r.Use(mid.AccessLog())
		r.Use(mid.Recovery())
	}

	//中间件
	r.Use(mid.Translations())
	r.Use(mid.Tracing())
	r.Use(mid.RateLimiter(methodLimiters))
	r.Use(mid.ContextTimeout(global.AppSetting.DefaultContextTimeout))
	r.Use(mid.Translations())
	r.Use(mid.Tracing())

	//文件服务
	upload := api.NewUpload()
	r.POST("upload/file", upload.UploadFile) //上传文件
	r.StaticFS("/static", http.Dir(global.AppSetting.UploadSavePath))

	//jwt
	r.GET("/auth", api.GetAuth)

	//业务接口
	article := v1.NewArticle()
	tag := v1.NewTag()
	apiv1 := r.Group("api/v1")
	apiv1.Use(mid.JWT())
	{
		apiv1.POST("/tags", tag.Create)
		apiv1.DELETE("/tags/:id", tag.Delete)
		apiv1.PUT("/tags/:id", tag.Update)   // put: 更新动作，用于更新一个完整的资源，要求为幂等
		apiv1.PATCH("/tags/:id", tag.Update) // patch: 更新动作，用于更新某一资源的某一个组成部分，可以不幂等
		apiv1.GET("/tags/:id", tag.Get)
		apiv1.GET("/tags", tag.List)

		apiv1.POST("/articles", article.Create)
		apiv1.DELETE("/articles/:id", article.Delete)
		apiv1.PUT("/articles/:id", article.Update)
		apiv1.PATCH("/articles/:id", article.Update)
		apiv1.GET("/articles/:id", article.Get)
		apiv1.GET("/articles", article.List)
	}

	return r
}
