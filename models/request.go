package models

type ReqGameList struct {
	Page  int `form:"page"`
	Limit int `form:"limit"`
}

type ReqLogin struct {
	Username string `form:"username"`
	Password string `form:"password"`
}

type ReqRegister struct {
	Username string `form:"username"`
	Password string `form:"password"`
	Email    string `form:"email"`
}
