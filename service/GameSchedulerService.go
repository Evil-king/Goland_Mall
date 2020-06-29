package service

import (
	"Goland_Mall/dao"
	"Goland_Mall/dto"
	"Goland_Mall/model"
	"Goland_Mall/serializer"
)

func GameSchedulerList() serializer.Result {
	gameSchedulerList := dao.GameSchedulerList()
	if len(gameSchedulerList)==0{
		return serializer.Fail(nil, nil)
	}
	return serializer.SuccessData(gameSchedulerList)
}

func GameSchedulerUpdate(gameSchedulerUpdate dto.GameSchedulerDto) serializer.Result {
	affected := dao.GameSchedulerUpdate(gameSchedulerUpdate)
	if affected > 0 {
		return serializer.Success()
	}
	return serializer.Fail(nil, nil)
}

func GetSchedulerByGameCode(gameCode string) model.GameScheduler  {
	gameScheduler :=dao.GetSchedulerByGameCode(gameCode)
	return gameScheduler
}
