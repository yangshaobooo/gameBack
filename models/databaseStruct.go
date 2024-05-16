package models

type Category struct {
	CategoryID   string `db:"categoryID"`
	CategoryName string `db:"categoryName"`
	GameCount    string `db:"gameCount"`
}
