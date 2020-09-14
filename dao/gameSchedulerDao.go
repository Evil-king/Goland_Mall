package dao

import (
	"Game/model"
	"Game/utils"
)

//游戏计划
func GameSchedulerList() []*model.GameScheduler {
	var gameSchedulerList []*model.GameScheduler
	db := utils.DbHelper
	db.Model(&model.GameScheduler{}).Find(&gameSchedulerList)
	return gameSchedulerList
}

//更新游戏计划
func GameSchedulerUpdate(gameScheduler model.GameScheduler) int64 {
	db := utils.DbHelper
	affected := db.Model(&model.GameScheduler{}).Where("game_code = ?", gameScheduler.GameCode).RowsAffected
	return affected
}

func GetSchedulerByGameCode(gameCode string) model.GameScheduler {
	db := utils.DbHelper
	var gameScheduler model.GameScheduler
	db.Model(&model.GameScheduler{}).Where("game_code = ?", gameCode).Find(&gameScheduler)
	return gameScheduler
}

func SelectGameSchedulerToEveryDay(gameCode string, modelCode string) model.GameScheduler {
	db := utils.DbHelper
	var gameScheduler model.GameScheduler
	db.Model(&model.GameScheduler{}).Where("game_code = ? and model_code = ? and draw_cycle = ?",
		gameCode, modelCode,"everyDay").
		Find(&gameScheduler)
	return gameScheduler
}

func SelectGameSchedulerToEveryWeek(gameCode string, modelCode string) model.GameScheduler {
	db := utils.DbHelper
	var gameScheduler model.GameScheduler
	db.Model(&model.GameScheduler{}).Where("game_code = ? and model_code = ? and draw_cycle = ?",
		gameCode, modelCode,"everyWeek" ).
		Find(&gameScheduler)
	return gameScheduler
}
