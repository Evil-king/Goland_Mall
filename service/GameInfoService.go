package service

import (
	"Game/common/dto"
	"Game/dao"
	"Game/model"
)

type IGameInfoService interface {
	GameInfoList(gameInfoDto dto.GameInfoDto) (model.Page, error)
	CreateGameInfo(gameInfo *model.GameInfo) string
	UpdateGameInfo(gameCode string, gameStatus string) string
}
//实现类
type GameInfoService struct {}

func (g *GameInfoService) GameInfoList(gameInfoDto dto.GameInfoDto) (model.Page, error) {
	gameInfos, err := dao.GameInfoList(gameInfoDto)
	return gameInfos,err
}

func (g *GameInfoService) CreateGameInfo(gameInfo *model.GameInfo) string {
	return dao.CreateGameInfo(gameInfo)
}

func (g *GameInfoService) UpdateGameInfo(gameCode string, gameStatus string) string {
	return dao.UpdateGameInfo(gameCode, gameStatus)
}


