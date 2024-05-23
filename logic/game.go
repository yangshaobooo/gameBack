package logic

import (
	"go.uber.org/zap"
	"webGameBack/dao/mysql"
	"webGameBack/models"
)

func GameListAll(req *models.ReqGameListAll) (*models.RespGame, error) {
	// 计算偏移量
	offset := (req.Page - 1) * req.Limit
	gameList, err := mysql.GameListAll(offset, req.Limit)
	if err != nil {
		zap.L().Error("logic GameListAll dao.GameListAll failed", zap.Error(err))
		return nil, err
	}
	res := &models.RespGame{
		Response: models.Response{
			StatusCode: 1,
			StatusMsg:  "获取全部游戏列表成功",
		},
		Games: gameList,
	}
	return res, nil
}
