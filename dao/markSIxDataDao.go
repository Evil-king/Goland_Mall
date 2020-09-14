package dao

import (
	"Game/model"
	"Game/utils"
)

func SelectByYear(year string) model.MarkSixData {
	var markSixData model.MarkSixData
	db := utils.DbHelper
	db.Model(&model.MarkSixData{}).Where("year=?", year).Find(&markSixData)
	return markSixData
}

func SelectAllYear() []*model.MarkSixData {
	var markSixData []*model.MarkSixData
	db := utils.DbHelper
	db.Model(&model.MarkSixData{}).Find(&markSixData)
	return markSixData
}