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
