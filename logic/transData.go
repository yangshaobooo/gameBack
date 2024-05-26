package logic

import (
	"encoding/json"
	"go.uber.org/zap"
	"webGameBack/models"
)

func ParseGameDetail(detail *models.Detail) *models.DetailParse {
	res := new(models.DetailParse)
	// json数据转化为数组
	var pics []string
	if err := json.Unmarshal([]byte(detail.Picture), &pics); err != nil {
		zap.L().Error("logic transData.go ParseGameDetail json.Unmarshal failed", zap.Error(err))
		return nil
	}
	var media []models.TitleGifText
	if err := json.Unmarshal([]byte(detail.TitleGif), &media); err != nil {
		zap.L().Error("logic transData.go ParseGameDetail media json.Unmarshal failed", zap.Error(err))
		return nil
	}
	res.GameID = detail.GameID
	res.Video = detail.Video
	res.Pictures = pics
	res.Description = detail.Description
	res.TitleGifs = media

	return res
}
