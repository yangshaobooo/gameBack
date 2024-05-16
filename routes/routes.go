package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"webGameBack/controller"
	"webGameBack/logger"
	"webGameBack/middlewares"
)

func Setup() *gin.Engine {
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true), middlewares.CorsMiddleware())
	r.GET("/", func(context *gin.Context) {
		context.String(http.StatusOK, "ok")
	})

	// 游戏路由
	apiRouter := r.Group("/webGame")

	// 游戏分类
	apiRouter.GET("/category/", controller.Category)

	return r
}
