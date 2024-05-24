package mysql

import (
	"errors"
	"go.uber.org/zap"
	"webGameBack/models"
)

func QueryGameCategory() ([]models.Category, error) {
	var res []models.Category
	sqlStr := `Select categoryID, categoryName, gameCount from game_category`
	if err := db.Select(&res, sqlStr); err != nil {
		zap.L().Error("mysql QueryGameCategory failed", zap.Error(err))
		return nil, errors.New("mysql 获取类别失败")
	}
	return res, nil
}
