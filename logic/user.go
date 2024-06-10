package logic

import (
	"go.uber.org/zap"
	"time"
	"webGameBack/dao/mysql"
	"webGameBack/middlewares/jwt"
	"webGameBack/models"
)

// Login 处理登录逻辑
func Login(p *models.ReqLogin) (*models.RespLogin, error) {
	// 1、查询数据库中用户信息
	user, err := mysql.Login(p)
	if err != nil {
		zap.L().Error("mysql.Login() failed", zap.Error(err))
		return nil, err
	}
	// 2、生成token
	token, err := jwt.GenToken(user.UserID, user.Username)
	if err != nil {
		return nil, err
	}
	// 3、判断是否具有会员权限
	currentTime := time.Now()
	var permit int8
	if user.Flag == 1 && currentTime.Before(user.EndTime) {
		permit = 1
	}
	// 4、组合返回数据
	res := &models.RespLogin{
		Response: models.Response{
			StatusCode: 1,
			StatusMsg:  "登录成功",
		},
		UserID:   user.UserID,
		Username: user.Username,
		Token:    token,
		Permit:   permit,
	}
	return res, nil
}
