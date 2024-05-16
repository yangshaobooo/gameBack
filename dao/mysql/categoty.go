package mysql

import (
	"errors"
	"go.uber.org/zap"
	"webGameBack/models"
)

func QueryGameCategory() ([]models.Category, error) {
	var res []models.Category
	// 编写查询语句
	sqlStr := `Select categoryID, categoryName, gameCount from game_category`
	if err := db.Select(&res, sqlStr); err != nil {
		zap.L().Error("mysql QueryGameCategory failed", zap.Error(err))
		return nil, errors.New("mysql 获取类别失败")
	}
	// 查询category数据成功
	return res, nil
}
