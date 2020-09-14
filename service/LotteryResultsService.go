package service

import (
	"Game/common/dto"
	"Game/dao"
	"Game/model"
)


type ILotteryResultService interface {
	GetLotteryResultsList(dto *dto.LotteryResultsDto) model.Page
}

type LotteryResultService struct {}

//获取开奖结果
func (l *LotteryResultService) GetLotteryResultsList(dto *dto.LotteryResultsDto) model.Page {
	return dao.GetLotteryResultsList(dto)
}

func (l *LotteryResultService) GetCurrentPeriod(gameCode string) model.LotteryResults  {
	return dao.GetCurrentPeriod(gameCode)
}
