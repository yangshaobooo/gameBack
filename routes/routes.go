package routes

import (
	"github.com/gin-gonic/gin"
	"webGameBack/controller"
	"webGameBack/logger"
	"webGameBack/middlewares"
)

func Setup() *gin.Engine {
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true), middlewares.CorsMiddleware())

	// 设置路由组
	apiRouter := r.Group("/webGame")

	apiRouter.GET("/category", controller.Category)             // 获取分类
	apiRouter.GET("/list", controller.GameListAll)              // 获取游戏
	apiRouter.GET("/list/:categoryID", controller.GameListPart) // 获取分类游戏
	apiRouter.GET("/detail/:gameID", controller.GameDetail)     // 获取游戏详情
	apiRouter.GET("/download/:gameID", controller.Download)     // 获取下载链接
	apiRouter.POST("/login", controller.Login)                  //登录
	apiRouter.POST("/register", controller.Register)            // 注册
	return r
}
