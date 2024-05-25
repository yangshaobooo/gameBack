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
