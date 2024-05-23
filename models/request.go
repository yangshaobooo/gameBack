package models

type ReqGameListAll struct {
	Page  int `form:"page"`
	Limit int `form:"limit"`
}
