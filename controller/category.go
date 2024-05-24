package controller

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"webGameBack/logic"
)

// Category 获取分类
func Category(c *gin.Context) {
	res, err := logic.GetCategory()
	if err != nil {
		zap.L().Error("logic.GetCategory() failed", zap.Error(err))
		return
	}
	// 返回数据
	c.JSON(http.StatusOK, res)
}
