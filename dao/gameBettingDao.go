package dao

import (
	"Game/common/dto"
	"Game/common/vo"
	"Game/model"
	"Game/utils"
	"github.com/shopspring/decimal"
)

var db = utils.DbHelper

//获取投注项列表
func GameDetails(gameCode string) []*vo.GameInfoDetailsVo {
	var gameInfoDetailsList []*vo.GameInfoDetailsVo
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
		affected = int(db.Model(&model.GameBetting{}).Where("id = ?", obj.BettingId).Update("betting_status", obj.Flag).RowsAffected)
	}
	if affected > 0 {
		return "SUCCESS"
	}
	return "FAIL"
}

func SelectGameInfoByGameCode(gameCode string, modelCode string) []*vo.GameInfoInnerAggregation {
	var gameInfoInnerAggregation []*vo.GameInfoInnerAggregation
	db.Table("game_betting").Select("game_betting.math_odds,game_betting.betting_name,game_betting.betting_status,"+
		"gg.group_name,gm.method_name,game_betting.id,gf.game_name,gf.game_code").
		Joins("LEFT JOIN game_items as gm ON gm.id = game_betting.items_id").
		Joins("LEFT JOIN game_group as gg ON gg.id=gm.group_id").
		Joins("LEFT JOIN game_info gf ON gf.model_code = gg.model_code").
		Joins("LEFT JOIN game_model m on m.model_code = gf.model_code").
		Where("gf.game_code = ? and m.model_code = ? and m.model_status = 'open'", gameCode, modelCode).
		Find(&gameInfoInnerAggregation)
	return gameInfoInnerAggregation
}

func UpdateGameBettingOddsByIdInZodiac(id int64, mathOdds decimal.Decimal, bettingName string) {
	affected := db.Model(model.GameBetting{}).
		Where("id = ? and betting_name = ?", id, bettingName).
		Update("math_odds = ?", mathOdds).RowsAffected
	if affected < 0 {
		panic("更新失败")
	}
}

func UpdateGameBettingOddsById(id int64, mathOdds decimal.Decimal) {
	affected := db.Model(model.GameBetting{}).
		Where("id = ?", id).
		Update("math_odds = ?", mathOdds).RowsAffected
	if affected < 0 {
		panic("更新失败")
	}
}

func InnerGameInfo(gameCode string, flag string) []*vo.GameInfoInnerVO {
	var gameInfoInnerList []*vo.GameInfoInnerVO
	db.Table("game_betting").Select("game_betting.math_odds,game_betting.betting_name,game_betting.betting_status,"+
		"gg.group_name,gm.method_name,game_betting.id,game_betting.attributes,game_betting.lottery_attributes," +
		"game_betting.page_attributes,gf.game_name,gf.game_code").
		Joins("LEFT JOIN game_items as gm ON gm.id = game_betting.items_id").
		Joins("LEFT JOIN game_group as gg ON gg.id=gm.group_id").
		Joins("LEFT JOIN game_info gf ON gf.model_code = gg.model_code").
		Joins("LEFT JOIN game_model m on m.model_code = gf.model_code").
		Where("gf.game_code = ? and game_betting.betting_status=?", gameCode,flag).
		Find(&gameInfoInnerList)
	return gameInfoInnerList
}
