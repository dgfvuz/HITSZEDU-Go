package router

import (
	"hitszedu-go/config"
	"hitszedu-go/controller/other"
	"hitszedu-go/controller/user"
	"hitszedu-go/middleware"

	"github.com/gin-gonic/gin"
)

func InitRouter() {
	gin.SetMode(gin.ReleaseMode)
	addr := ":" + config.GetString("server.port")
	router := gin.Default()
	//使用中间件
	router.Use(middleware.Cors())
	//加载静态资源
	// router.LoadHTMLFiles("./static/index.html")
	// router.Static("/css", "./static/css")
	// router.Static("/js", "./static/js")
	// router.StaticFile("/favicon.ico", "./static/favicon.ico")
	//请求

	userGroup := router.Group("/api/user")
	{
		userGroup.POST("/wxlogin", user.Wxlogin)
		userGroup.POST("/wxregister", user.WxRegister)
	}

	otherGroup := router.Group("/api/other")
	{
		otherGroup.Use(middleware.JWTAuth())
		otherGroup.POST("/addadvice", other.AddAdvice)
	}

	router.Run(addr)
}
