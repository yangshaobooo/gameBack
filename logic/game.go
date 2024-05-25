package logic

import (
	"go.uber.org/zap"
	"webGameBack/dao/mysql"
	"webGameBack/models"
)

func GameListAll(req *models.ReqGameList) (*models.RespGame, error) {
	// 计算偏移量
	offset := (req.Page - 1) * req.Limit
	gameList, err := mysql.GameListAll(offset, req.Limit)
	if err != nil {
		zap.L().Error("logic GameListAll mysql.GameListAll failed", zap.Error(err))
		return nil, err
	}
	gameCount, err := mysql.GameCount()
	if err != nil {
		zap.L().Error("logic GameListAll mysql.GameCount failed", zap.Error(err))
		return nil, err
	}
	res := &models.RespGame{
		Response: models.Response{
			StatusCode: 1,
			StatusMsg:  "获取全部游戏成功",
		},
		Games:        gameList,
		GameCategory: "首页：全部游戏",
		Total:        gameCount,
	}
	return res, nil
}

func GameListPart(req *models.ReqGameList, categoryID string) (*models.RespGame, error) {
	// 计算偏移量
	offset := (req.Page - 1) * req.Limit
	gameList, err := mysql.GameListPart(offset, req.Limit, categoryID)
	if err != nil {
		zap.L().Error("logic GameListPart mysql.GameListPart failed", zap.Error(err))
		return nil, err
	}
	gameCount, err := mysql.GameCountPart(categoryID)
	if err != nil {
		zap.L().Error("logic GameListPart mysql.GameCountPart failed", zap.Error(err))
		return nil, err
	}
	categoryStr := "没有该分类游戏"
	if len(gameList) != 0 {
		categoryStr = gameList[0].CategoryName
	}
	res := &models.RespGame{
		Response: models.Response{
			StatusCode: 1,
			StatusMsg:  "获取分类游戏成功",
		},
		Games:        gameList,
		GameCategory: categoryStr,
		Total:        gameCount,
	}
	return res, nil
}

// GameDetail 获取游戏详情
func GameDetail(gameID string) (*models.RespDetail, error) {
	// todo
	// 查找该游戏的视频图片和介绍内容
	detail, err := mysql.GameDetail(gameID)
	if err != nil {
		zap.L().Error("logic GameDetail mysql.GameDetail() failed", zap.Error(err))
		return nil, err
	}
	game, err := mysql.GetGameByID(gameID)
	if err != nil {
		zap.L().Error("logic GameDetail mysql.GetGameByID() failed", zap.Error(err))
		return nil, err
	}
	res := &models.RespDetail{
		Response: models.Response{
			StatusCode: 1,
			StatusMsg:  "获取游戏详情成功",
		},
		Game:   *game,
		Detail: *detail,
	}
	return res, nil
}
