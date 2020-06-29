package dao

import (
	"Goland_Mall/dto"
	"Goland_Mall/model"
	"Goland_Mall/utils"
)

//游戏计划
func GameSchedulerList() []*model.GameScheduler {
	var gameSchedulerList []*model.GameScheduler
	db := utils.DbHelper
	db.Model(&model.GameScheduler{}).Find(&gameSchedulerList)
	return gameSchedulerList
}

//更新游戏计划
func GameSchedulerUpdate(gameSchedulerUpdate dto.GameSchedulerDto)  int64{
	db := utils.DbHelper
	affected := db.Model(&model.GameScheduler{}).Where("game_code = ?",gameSchedulerUpdate.GameCode).
		Update("draw_stime",gameSchedulerUpdate.DrawStartTime).
		Update("draw_etime",gameSchedulerUpdate.DrawEndTime).
		Update("overall_time",gameSchedulerUpdate.OverallTime).
		Update("seal_time",gameSchedulerUpdate.SealTime).RowsAffected
	return affected
}

func GetSchedulerByGameCode(gameCode string) model.GameScheduler  {
	db := utils.DbHelper
	var gameScheduler model.GameScheduler
	db.Model(&model.GameScheduler{}).Where("game_code = ?",gameCode).Find(&gameScheduler)
	return gameScheduler
}