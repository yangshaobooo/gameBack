package controller

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"webGameBack/logic"
	"webGameBack/models"
)

// GameListAll 获取游戏列表
func GameListAll(c *gin.Context) {
	// todo 进行深分页优化
	// 获取分页参数
	requestData := new(models.ReqGameList)
	if err := c.ShouldBindQuery(requestData); err != nil {
		zap.L().Error("controller.game get requestData failed", zap.Error(err))
		return
	}

	// 获取全部游戏列表返回数据
	res, err := logic.GameListAll(requestData)
	if err != nil {
		zap.L().Error("controller GameListAll() logic.GameListAll() failed", zap.Error(err))
		return
	}
	c.JSON(http.StatusOK, res)
}

// GameListPart 获取分类游戏列表
func GameListPart(c *gin.Context) {
	// todo 进行分页优化
	// 获取参数
	categoryID := c.Param("categoryID")
	requestData := new(models.ReqGameList)
	if err := c.ShouldBindQuery(requestData); err != nil {
		zap.L().Error("controller GameLIstPart  get requestData failed", zap.Error(err))
		return
	}

	//获取游戏数据
	res, err := logic.GameListPart(requestData, categoryID)
	if err != nil {
		zap.L().Error("controller GameListPart() failed", zap.Error(err))
	}
	c.JSON(http.StatusOK, res)
}
