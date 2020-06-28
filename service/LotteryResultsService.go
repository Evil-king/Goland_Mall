package service

import (
	"Goland_Mall/dao"
	"Goland_Mall/model"
	"Goland_Mall/serializer"
)

func GetLotteryResultsList(dto *model.LotteryResultsDto) serializer.Result  {
	result := dao.GetLotteryResultsList(dto)
	return serializer.Success(result,nil)
}
