package dao

import (
	"Goland_Mall/dto"
	"Goland_Mall/model"
	"Goland_Mall/utils"
	"Goland_Mall/vo"
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
	affected := db.Where("game_code = ?", gameCode).Update("game_status", gameStatus).RowsAffected
	if affected > 0 {
		return "SUCCESS"
	}
	return "FAIL"
}

//获取投注项列表
func GameDetails(gameCode string) []*vo.GameInfoDetailsVo {
	var gameInfoDetailsList []*vo.GameInfoDetailsVo
	db := utils.DbHelper
	//gorm联表查询
	db.Table("game_betting").Select("game_betting.math_odds,game_betting.betting_name,game_betting.betting_status,gg.group_name,gm.method_name,game_betting.id").
		Joins("LEFT JOIN game_items as gm ON gm.id = game_betting.items_id").
		Joins("LEFT JOIN game_group as gg ON gg.id=gm.group_id").
		Joins("LEFT JOIN game_info gf ON gf.model_code = gg.model_code").
		Joins("LEFT JOIN game_model m on m.model_code = gf.model_code").
		Where("gf.game_code = ? and m.model_status = 'open'", gameCode).Find(&gameInfoDetailsList)
	//原生联表查询
	//db.Raw("SELECT gg.group_name,gm.method_name,gb.betting_name,gb.math_odds,gb.betting_status,gb.id " +
	//	"FROM game_betting gb LEFT JOIN game_items gm ON gm.id = gb.items_id LEFT JOIN game_group gg ON gg.id=gm.group_id LEFT JOIN game_info gf " +
	//	"ON gf.model_code = gg.model_code LEFT JOIN game_model m on m.model_code = gf.model_code where gf.game_code = ? and m.model_status = 'open'",gameCode).Scan(&gameInfoDetailsList)
	return gameInfoDetailsList
}

//投注项开关
func MathOddsFlag(params []*dto.BettingMathOddsFlgDto) string {
	var affected int
	//遍历切片
	for i := 0; i < len(params); i++ {
		obj := params[i]
		db := utils.DbHelper
		affected = int(db.Model(&model.GameBetting{}).Where("id = ?", obj.BettingId).Update("betting_status", obj.Flag).RowsAffected)
	}
	if affected > 0 {
		return "SUCCESS"
	}
	return "FAIL"
}

func SelectGameInfoByGameCode(gameCode string) []*vo.GameInfoInnerAggregation {
	var gameInfoInnerAggregation []*vo.GameInfoInnerAggregation
	db := utils.DbHelper
	db.Table("game_betting").Select("game_betting.math_odds,game_betting.betting_name,game_betting.betting_status,"+
		"gg.group_name,gm.method_name,game_betting.id,gf.game_name,gf.game_code").
		Joins("LEFT JOIN game_items as gm ON gm.id = game_betting.items_id").
		Joins("LEFT JOIN game_group as gg ON gg.id=gm.group_id").
		Joins("LEFT JOIN game_info gf ON gf.model_code = gg.model_code").
		Joins("LEFT JOIN game_model m on m.model_code = gf.model_code").
		Where("gf.game_code = ? and m.model_status = 'open'", gameCode).Find(&gameInfoInnerAggregation)

	return gameInfoInnerAggregation
}
