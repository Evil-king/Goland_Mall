package service

import (
	"Game/common/vo"
	"Game/dao"
	"Game/model"
)

type ZodiacOddsRelationService struct {
}

type IZodiacOddsRelationService interface {
	SelectByItemsName(itemName string) model.ZodiacOddsRelation
	SelectByItemsNameBettingName(itemName string, bettingName string) model.ZodiacOddsRelation
	GetZodiacList() []vo.ZodiacVo
}

func (z *ZodiacOddsRelationService) SelectByItemsName(itemName string) model.ZodiacOddsRelation {
	return dao.SelectByItemsName(itemName)
}

func (z *ZodiacOddsRelationService) SelectByItemsNameBettingName(itemName string, bettingName string) model.ZodiacOddsRelation {
	return dao.SelectByItemsNameBettingName(itemName,bettingName)
}

func (z *ZodiacOddsRelationService) GetZodiacList() []vo.ZodiacVo  {
	return dao.GetZodiacList()
}
