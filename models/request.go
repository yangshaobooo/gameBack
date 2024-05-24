package models

type ReqGameList struct {
	Page  int `form:"page"`
	Limit int `form:"limit"`
}
