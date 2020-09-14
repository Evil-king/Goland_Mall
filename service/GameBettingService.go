package service

import (
	"Game/common/dto"
	"Game/common/vo"
	"Game/dao"
	"github.com/shopspring/decimal"
)

type GameBettingService struct {
}

type IGameBettingService interface {
	GameDetails(gameCode string) []*vo.GameInfoDetailsVo
	MathOddsFlag(params []*dto.BettingMathOddsFlgDto) string
	UpdateGameBettingOddsByIdInZodiac(id int64, mathOdds string, bettingName string)
	UpdateGameBettingOddsById(id int64, mathOdds string)
	InnerGameInfo(gameCode string, flag string) []*vo.GameInfoInnerVO
}

func (g *GameBettingService) GameDetails(gameCode string) []*vo.GameInfoDetailsVo {
	return dao.GameDetails(gameCode)
}

func (g *GameBettingService) MathOddsFlag(params []*dto.BettingMathOddsFlgDto) string {
	return dao.MathOddsFlag(params)
}

func (g *GameBettingService) UpdateGameBettingOddsByIdInZodiac(id int64, mathOdds decimal.Decimal, bettingName string) {
	dao.UpdateGameBettingOddsByIdInZodiac(id, mathOdds, bettingName)
}

func (g *GameBettingService) UpdateGameBettingOddsById(id int64, mathOdds decimal.Decimal) {
	dao.UpdateGameBettingOddsById(id, mathOdds)
}

func (g *GameBettingService) InnerGameInfo(gameCode string, flag string) []*vo.GameInfoInnerVO {
	if flag == ""{
		flag = "1"
	}
	return dao.InnerGameInfo(gameCode, flag)
}
