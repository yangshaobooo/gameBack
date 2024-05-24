package mysql

import (
	"go.uber.org/zap"
	"webGameBack/models"
)

func GameListAll(offset int, limit int) ([]models.Game, error) {
	var gameList []models.Game
	str := `select gameID,gameName,categoryID,category,imageUrl,updateTime,watching,price,size,introduce from game_list LIMIT ? OFFSET ?`
	if err := db.Select(&gameList, str, limit, offset); err != nil {
		zap.L().Error("dao mysql game db.Select() failed", zap.Error(err))
		return gameList, err
	}
	return gameList, nil
}

func GameListPart(offset int, limit int, categoryID string) ([]models.Game, error) {
	var gameList []models.Game
	str := `select gameID,gameName,categoryID,category,imageUrl,updateTime,watching,price,size,introduce from game_list where categoryID=? LIMIT ? OFFSET ?`
	if err := db.Select(&gameList, str, categoryID, limit, offset); err != nil {
		zap.L().Error("dao GameListPart sb.Select() failed", zap.Error(err))
		return gameList, err
	}
	return gameList, nil
}
