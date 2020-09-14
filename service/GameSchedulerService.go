package service

import (
	"Game/common/dto"
	"Game/dao"
	"Game/model"
	"Game/utils"
	"time"
)

type GameSchedulerService struct {
}

type IGameSchedulerService interface {
	GameSchedulerList() []*model.GameScheduler
	GameSchedulerUpdate(gameSchedulerUpdate dto.GameSchedulerDto) bool
	GetSchedulerByGameCode(gameCode string) model.GameScheduler
	SelectGameSchedulerToEveryDay(gameCode string, modelCode string) model.GameScheduler
	SelectGameSchedulerToEveryWeek(gameCode string, modelCode string) model.GameScheduler
}

func (g *GameSchedulerService) GameSchedulerList() []*model.GameScheduler {
	return dao.GameSchedulerList()
}

func (g *GameSchedulerService) GameSchedulerUpdate(gameSchedulerUpdate dto.GameSchedulerDto) bool {
	gameScheduler := g.GetSchedulerByGameCode(gameSchedulerUpdate.GameCode)
	var lotteryResultService = &LotteryResultService{}
	lotteryResultService.GetCurrentPeriod(gameScheduler.GameCode)

	if "everyDay" == gameSchedulerUpdate.DrawCycle {
		//将整体时常放入缓存中
		conn := utils.RedisDefaultPool.Get()
		_, err := conn.Do("HSET",
			"mp",
			"sealTime", gameScheduler.SealTime,
			"betTime", gameScheduler.BetTime)
		_, err = conn.Do("expire", gameSchedulerUpdate.GameCode+gameSchedulerUpdate.ModelCode, gameScheduler.BetTime+10) //10秒过期
		if err != nil {
			return false
		}
		gameScheduler := model.GameScheduler{
			Id:          utils.IdWork(),
			GameName:    "",
			GameCode:    gameSchedulerUpdate.GameCode,
			ModelCode:   gameSchedulerUpdate.ModelCode,
			DrawDay:     "",
			DrawStime:   gameSchedulerUpdate.DrawStartTime,
			DrawEtime:   gameSchedulerUpdate.DrawEndTime,
			OverAllTime: gameSchedulerUpdate.BetTime + 10,
			BetTime:     gameSchedulerUpdate.BetTime,
			ModifyTime:  time.Time{},
		}
		affected := dao.GameSchedulerUpdate(gameScheduler)
		if affected > 0 {
			return true
		}
	} else if "everyWeek" == gameSchedulerUpdate.DrawCycle {
		//affected := dao.GameSchedulerUpdate(gameSchedulerUpdate)
		//if affected > 0 {
		//	return true
		//}
	}
	return false
}

func (g *GameSchedulerService) GetSchedulerByGameCode(gameCode string) model.GameScheduler {
	gameScheduler := dao.GetSchedulerByGameCode(gameCode)
	return gameScheduler
}

func (g *GameSchedulerService) SelectGameSchedulerToEveryDay(gameCode string, modelCode string) model.GameScheduler {
	return dao.SelectGameSchedulerToEveryDay(gameCode, modelCode)
}

func (g *GameSchedulerService) SelectGameSchedulerToEveryWeek(gameCode string, modelCode string) model.GameScheduler {
	return dao.SelectGameSchedulerToEveryWeek(gameCode, modelCode)
}
