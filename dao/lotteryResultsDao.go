package dao

import (
	"Goland_Mall/model"
	"Goland_Mall/utils"
	"Goland_Mall/vo"
	"time"
)

func GetLotteryResultsList(dto *model.LotteryResultsDto) model.Page {
	var lotteryResultsList []*model.LotteryResults
	var lotteryResultsListVo []*vo.LotteryResultsListVo
	//设置一个变量接收总记录数
	var totalRecord int64
	db := utils.DbHelper
	if dto.PeriodNum != "" {
		db = db.Where("period_num = ?", dto.PeriodNum)
	}
	if dto.StartTime != "" && dto.EndTime != "" {
		db = db.Where("create_time BETWEEN ? AND ?", dto.StartTime, dto.EndTime)
	}
	db.Limit(dto.PageSize).Offset((dto.CurrentPage - 1) * dto.PageSize).Find(&lotteryResultsList)
	//获取总数
	db.Model(model.LotteryResults{}).Count(&totalRecord)

	for i := 0; i < len(lotteryResultsList); i++ {
		lotteryResults := vo.LotteryResultsListVo{
			PeriodNum:      lotteryResultsList[i].PeriodNum,
			DrawTime:       &lotteryResultsList[i].DrawTime,
			WinningResults: lotteryResultsList[i].WinningResults,
			OutNumber:      lotteryResultsList[i].OutNumber,
		}
		lotteryResultsListVo = append(lotteryResultsListVo, &lotteryResults)
	}
	return model.OperatorData(lotteryResultsListVo, totalRecord, dto.CurrentPage, dto.PageSize)
}

func SelectPeriodNumByStatus(gameCode string) model.LotteryResults {
	var lotteryResults model.LotteryResults
	db := utils.DbHelper
	db.Model(&model.LotteryResults{}).Where("game_code = ?", gameCode).Find(&lotteryResults)
	return lotteryResults
}

func SelectPeriodNumByIsClose(gameCode string) model.LotteryResults {
	var lotteryResults model.LotteryResults
	db := utils.DbHelper
	db.Model(&model.LotteryResults{}).Where("game_code = ? and status =?", gameCode, "waiting").Order("draw_time desc").Limit(1).Find(&lotteryResults)
	return lotteryResults
}

func InsertData(gameCode string, periodNum string, outNumber []int, winningResults []string, flag string) {
	lotteryResults := &model.LotteryResults{
		Id:             utils.IdWork(),
		GameCode:       gameCode,
		PeriodNum:      periodNum,
		DrawTime:       time.Time{},
		OutNumber:      utils.SliceToString(outNumber),
		WinningResults: utils.SliceToString(winningResults),
		IsClose:        flag,
		Status:         "waiting",
		CreateTime:     time.Now(),
		ModifyTime:     time.Now(),
	}
	utils.DbHelper.Create(&lotteryResults)
}

func SelectLastIssue(periodNum string) model.LotteryResults {
	var lotteryResults model.LotteryResults
	db := utils.DbHelper
	db.Model(&model.LotteryResults{}).Where("period_num = ?", periodNum).Find(&lotteryResults)
	return lotteryResults
}

func UpdateLotteryByParams(results model.LotteryResults) int {
	//更新
	db := utils.DbHelper
	affected := db.Model(model.LotteryResults{}).Where("id = ?", results.Id).Update(results).RowsAffected
	return int(affected)
}
