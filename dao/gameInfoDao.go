package dao

import (
	"Game/common/dto"
	"Game/model"
	"Game/utils"
)

//获取游戏列表数据
func GameInfoList(dto dto.GameInfoDto) (model.Page, error) {
	//基于gorm的写法
	var gameInfoList []*model.GameInfo

	db := utils.DbHelper

	//设置一个变量接收总记录数
	var totalRecord int64

	//全部查询并且带条件
	if dto.GameName != "" {
		db = db.Where("game_name LIKE ?", "%"+dto.GameName+"%")
	}
	if dto.ModelCode != "" {
		db = db.Where("model_code LIKE ?", "%"+dto.ModelCode+"%")
	}
	if dto.GameStatus != "" {
		db = db.Where("game_status = ?", dto.GameStatus)
	}
	db.Limit(dto.PageSize).Offset((dto.CurrentPage - 1) * dto.PageSize).Find(&gameInfoList)
	//获取总数
	db.Model(model.GameInfo{}).Count(&totalRecord)

	return model.OperatorData(gameInfoList, totalRecord, dto.CurrentPage, dto.PageSize), nil
}

//新增游戏
func CreateGameInfo(gameInfo *model.GameInfo) string {
	affected := utils.DbHelper.Create(&gameInfo).RowsAffected
	if affected > 0 {
		return "SUCCESS"
	}
	return "FAIL"
}

//修改游戏状态
func UpdateGameInfo(gameCode string, gameStatus string) string {
	db := utils.DbHelper
	affected := db.Model(model.GameInfo{}).
		Where("game_code = ?", gameCode).
		Update("game_status", gameStatus).RowsAffected
	if affected > 0 {
		return "SUCCESS"
	}
	return "FAIL"
}
