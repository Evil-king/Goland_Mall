package service

import (
	"Game/dao"
	"Game/model"
)

type MarkSixDataService struct {
}

type IMarkSixDataService interface {
	SelectByYear(year string)  model.MarkSixData
	SelectAllYear()  []*model.MarkSixData
}

func (m *MarkSixDataService) SelectByYear(year string)  model.MarkSixData {
	return dao.SelectByYear(year)
}

func (m *MarkSixDataService) SelectAllYear() []*model.MarkSixData {
	return dao.SelectAllYear()
}