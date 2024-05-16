package logic

import (
	"go.uber.org/zap"
	"log"
	"webGameBack/dao/mysql"
	"webGameBack/models"
)

func GetCategory() (*models.RespCategory, error) {
	log.Println("get gameCategory running.....")
	gameCategory, err := mysql.QueryGameCategory()
	if err != nil {
		zap.L().Error("GetCategory中 mysql.QueryGameCategory() failed ", zap.Error(err))
		return nil, err
	}
	res := &models.RespCategory{
		Response: models.Response{
			StatusCode: 1,
			StatusMsg:  "获取分类成功",
		},
		Categories: gameCategory,
	}
	return res, nil
}
