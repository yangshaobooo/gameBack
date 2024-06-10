package mysql

import (
	"database/sql"
	"errors"
	"webGameBack/models"
)

// Login 查询用户账号密码时候正确
func Login(p *models.ReqLogin) (*models.User, error) {
	user := new(models.User)
	// 1、查询用户是否存在
	strSql := `select user_id,username,password,email,flag,endTime from user where username=?`
	if err := db.Get(user, strSql, p.Username); err != nil {
		if errors.Is(err, sql.ErrNoRows) { // 没查到用户
			return nil, errors.New("该用户不存在")
		}
	}
	// 2、查到用户之后对比密码
	// todo 加密密码
	if p.Password != user.Password {
		return nil, errors.New("密码错误")
	}
	// 3、返回用户信息
	return user, nil
}
