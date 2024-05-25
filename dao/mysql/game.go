package mysql

import (
	"go.uber.org/zap"
	"webGameBack/models"
)

// GameListAll 获取游戏数据全部展示
func GameListAll(offset int, limit int) ([]models.Game, error) {
	var gameList []models.Game
	str := `select gameID,gameName,categoryID,category,imageUrl,updateTime,watching,price,size,introduce from game_list LIMIT ? OFFSET ?`
	if err := db.Select(&gameList, str, limit, offset); err != nil {
		zap.L().Error("dao mysql game db.Select() failed", zap.Error(err))
		return gameList, err
	}
	return gameList, nil
}

// GameListPart 获取游戏数据，分类展示
func GameListPart(offset int, limit int, categoryID string) ([]models.Game, error) {
	var gameList []models.Game
	str := `select gameID,gameName,categoryID,category,imageUrl,updateTime,watching,price,size,introduce from game_list where categoryID=? LIMIT ? OFFSET ?`
	if err := db.Select(&gameList, str, categoryID, limit, offset); err != nil {
		zap.L().Error("dao GameListPart sb.Select() failed", zap.Error(err))
		return gameList, err
	}
	return gameList, nil
}

// GameCount 获取所有游戏数量
func GameCount() (int, error) {
	var count int
	str := `select count(*) from game_list`
	if err := db.Get(&count, str); err != nil {
		zap.L().Error("dao GameCount db.Get() failed", zap.Error(err))
		return 0, err
	}
	return count, nil
}

// GameCountPart 获取分类游戏数量
func GameCountPart(categoryID string) (int, error) {
	var count int
	str := `select count(*) from game_list where categoryID=?`
	if err := db.Get(&count, str, categoryID); err != nil {
		zap.L().Error("dao GameCount db.Get() failed", zap.Error(err))
		return 0, err
	}
	return count, nil
}

// GameDetail 获取游戏详情
func GameDetail(gameID string) (*models.Detail, error) {
	detail := new(models.Detail)
	str := `select gameID,video,picture,description,titleGif from game_detail where gameID=?`
	if err := db.Get(detail, str, gameID); err != nil {
		zap.L().Error("dao GameDetail db.Get() failed", zap.Error(err))
		return nil, err
	}
	return detail, nil
}

// GetGameByID 根据id获取游戏的封面分类等信息
func GetGameByID(gameID string) (*models.Game, error) {
	game := new(models.Game)
	str := `select gameID,gameName,categoryID,category,imageUrl,updateTime,watching,price,size,introduce from game_list where gameID=?`
	if err := db.Get(game, str, gameID); err != nil {
		zap.L().Error("dao GetGameByID db.Get() failed", zap.Error(err))
		return nil, err
	}
	return game, nil
}
