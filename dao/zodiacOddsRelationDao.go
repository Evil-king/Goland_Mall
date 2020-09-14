package dao

import (
	"Game/common/vo"
	"Game/model"
	"Game/utils"
)

func SelectByItemsName(itemsName string) model.ZodiacOddsRelation {
	var zodiacOddsRelation model.ZodiacOddsRelation
	db := utils.DbHelper
	db.Model(model.ZodiacOddsRelation{}).Where("itemsName = ?", itemsName).Find(&zodiacOddsRelation)
	return zodiacOddsRelation
}

func SelectByItemsNameBettingName(itemsName string, bettingName string) model.ZodiacOddsRelation {
	var zodiacOddsRelation model.ZodiacOddsRelation
	db := utils.DbHelper
	db.Model(model.ZodiacOddsRelation{}).Where("items_name = ? and betting_name=?", itemsName, bettingName).
		Find(&zodiacOddsRelation)
	return zodiacOddsRelation
}

func GetZodiacList() []vo.ZodiacVo {
	var zodiacVoList []vo.ZodiacVo
	db := utils.DbHelper
	db.Model(model.GameBetting{}).Select("gb.id,gb.items_id itemsId,gb.betting_name bettingName,gi.method_name itemName").
		Joins("left join game_items gi on gb.items_id = gi.id").
		Where("gi.method_name in","特肖","正肖","一肖","连肖").Find(&zodiacVoList)
	return zodiacVoList
}
