package mysql

import (
	"go.uber.org/zap"
	"webGameBack/models"
)

func GameListAll(offset int, limit int) ([]models.Game, error) {
	var gameList []models.Game
	str := `select gameID,gameName,categoryID,imageUrl,updateTime,watching,price,size,introduce from game_list LiMIT ? OFFSET ?`
	if err := db.Select(&gameList, str, limit, offset); err != nil {
		zap.L().Error("dao mysql game db.Select() failed", zap.Error(err))
		return gameList, err
	}
	return gameList, nil
}
