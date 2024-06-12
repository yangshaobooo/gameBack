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
	// 1、获取前端传递的username和password
	p := new(models.ReqLogin)
	resp := new(models.Response)
	if err := c.ShouldBind(p); err != nil { // 不同的参数发送方式，使用不同的绑定方式
		zap.L().Error("Login with invalid param", zap.Error(err))
		c.JSON(400, resp)
		return
	}
	fmt.Println("登录的用户", p)
	// 2、处理登录的逻辑
	res, err := logic.Login(p)
	if err != nil {
		zap.L().Error("logic.Login() error", zap.Error(err))
		resp.StatusMsg = "账号密码错误"
		c.JSON(400, resp)
		return
	} else {
		zap.L().Info("用户登录成功", zap.String("username", p.Username))
	}
	// 3、返回响应
	c.JSON(http.StatusOK, res)
}

// Register 注册账号
func Register(c *gin.Context) {
	// 1、获取前端传递的数据
	p := new(models.ReqRegister)
	if err := c.ShouldBind(p); err != nil {
		zap.L().Error("Register ShouldBind() failed", zap.Error(err))
		return
	}
	// 2、处理注册逻辑
	resp, err := logic.Register(p)
	if err != nil {
		zap.L().Error("Register logic.Register() failed", zap.Error(err))
		c.JSON(400, *resp)
		return
	}
	// 3、返回正确数据
	c.JSON(http.StatusOK, resp)
}
