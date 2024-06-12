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

// CheckUserExist 判断用户是否存在
func CheckUserExist(username string) error {
	sqlStr := `select count(username) from user where username=?`
	var count int
	if err := db.Get(&count, sqlStr, username); err != nil {
		return err
	}
	if count > 0 {
		return errors.New("账号已经存在")
	}
	return nil
}

// InsertUser 插入用户信息
func InsertUser(p *models.ReqRegister, userID int64) error {
	sqlStr := `insert into user (user_id,username,password,email,flag)values(?,?,?,?,0)`
	_, err := db.Exec(sqlStr, userID, p.Username, p.Password, p.Email)
	if err != nil {
		return err
	}
	return nil
}
