package rpcservices

import (
	."Game/service"
	"context"
	"google.golang.org/protobuf/runtime/protoimpl"
)
var gameBettingService = &GameBettingService{}

type InnerGameInfoInterface struct {}

func (*InnerGameInfoInterface) InnerGameInfo(ctx context.Context,  request *InnerGameInfoRequest) (*InnerGameInfoResponse, error) {
	var gameInfoInnerVoList []*GameInfoInnerVo
	gameCode := request.GameCode
	flag := request.Flag
	list := gameBettingService.InnerGameInfo(gameCode,flag)
	for _,value := range list{
		gameInfoInnerVo := &GameInfoInnerVo{
			state:             protoimpl.MessageState{},
			sizeCache:         0,
			unknownFields:     nil,
			GameCode:          value.GameCode,
			GameName:          value.GameName,
			GroupName:         value.GroupName,
			MathOdds:          value.MathOdds,
			BettingName:       value.BettingName,
			BettingStatus:     value.BettingStatus,
			BettingId:         value.Id,
			MethodName:        value.MethodName,
			Attributes:        value.Attributes,
			LotteryAttributes: value.LotteryAttributes,
			PageAttributes:    value.PageAttributes,
			Sort:              value.Sort,
		}
		gameInfoInnerVoList = append(gameInfoInnerVoList,gameInfoInnerVo)
	}
	return &InnerGameInfoResponse{GameInfoInnerVo:gameInfoInnerVoList},nil
}
