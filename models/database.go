package models

import (
	"encoding/json"
	"time"
)

type Category struct {
	CategoryID   string `db:"categoryID"`
	CategoryName string `db:"categoryName"`
	GameCount    string `db:"gameCount"`
}

type Game struct {
	GameID       string    `db:"gameID"`
	GameName     string    `db:"gameName"`
	CategoryID   string    `db:"categoryID"`
	CategoryName string    `db:"category"`
	ImageUrl     string    `db:"imageUrl"`
	UpdateTime   time.Time `db:"updateTime"`
	Watching     uint32    `db:"watching"`
	Price        uint8     `db:"price"`
	Size         string    `db:"size"`
	Introduce    string    `db:"introduce"`
}

type Detail struct {
	GameID      string          `db:"gameID"`
	Video       string          `db:"video"`
	Picture     json.RawMessage `db:"picture"`
	Description string          `db:"description"`
	TitleGif    json.RawMessage `db:"titleGif"`
}

type Download struct {
	GameID   string `db:"gameID"`
	Link     string `db:"link"`
	Source   string `db:"source"`
	Position int    `db:"position"`
}

type User struct {
	UserID   int64     `db:"user_id"`
	Username string    `db:"username"`
	Password string    `db:"password"`
	Email    string    `db:"email"`
	Flag     int8      `db:"flag"`    // 是否是会员
	EndTime  time.Time `db:"endTime"` // 会员到期时间
}
