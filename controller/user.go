package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"webGameBack/logic"
	"webGameBack/models"
)

// Login 用户登录函数
func Login(c *gin.Context) {
	// 1、获取前端传递的account和password
	p := new(models.ReqLogin)
	if err := c.ShouldBind(p); err != nil { // 不同的参数发送方式，使用不同的绑定方式
		zap.L().Error("Login with invalid param", zap.Error(err))
		return
	}
	fmt.Println("登录的用户", p)
	// 2、处理登录的逻辑
	res, err := logic.Login(p)
	if err != nil {
		zap.L().Error("logic.Login() error", zap.Error(err))
		return
	} else {
		zap.L().Info("用户登录成功", zap.String("account", p.Username))
	}
	// 3、返回响应
	c.JSON(http.StatusOK, res)
}