package logic

import (
	"go.uber.org/zap"
	"time"
	"webGameBack/dao/mysql"
	"webGameBack/middlewares/jwt"
	"webGameBack/middlewares/snowflake1"
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

// Register 处理注册逻辑
func Register(p *models.ReqRegister) (*models.Response, error) {
	resp := &models.Response{StatusCode: 0, StatusMsg: "注册失败"}
	// 1、判断用户名是否在数据库中已经存在
	if err := mysql.CheckUserExist(p.Username); err != nil {
		zap.L().Error("用户已经存在", zap.Error(err))
		resp.StatusMsg = "用户名已经存在"
		return resp, err
	}
	// 2、为新用户生成唯一id
	userID := snowflake1.GenID()

	// 3、用户数据插入数据库
	if err := mysql.InsertUser(p, userID); err != nil {
		zap.L().Error("Register mysql.InsertUser() failed", zap.Error(err))
		return resp, err
	}
	// 4、返回注册成功
	resp.StatusCode = 1
	resp.StatusMsg = "注册成功"
	return resp, nil
}
