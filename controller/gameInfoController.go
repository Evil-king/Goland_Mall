package controller

import (
	"Game/common/dto"
	"Game/common/serializer"
	"Game/model"
	. "Game/service"
	"Game/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

var gameInfoService = &GameInfoService{}
var gameBettingService = &GameBettingService{}

//获取游戏列表
// @Summary 游戏列表
// @Produce json
// @Param GameName query string true "GameName"
// @Param GameName query string true "GameName"
// @Param GameName query string true "GameName"
// @Router /game/gameInfo [POST]
func PostGameInfoList(c *gin.Context) {
	var gameInfoDto dto.GameInfoDto
	error := c.BindJSON(&gameInfoDto)
	if error != nil {
		fmt.Printf("参数异常:%s", error.Error())
	}
	gameInfos, err := gameInfoService.GameInfoList(gameInfoDto)
	if err != nil {
		serializer.Fail(c, nil, "gameInfoList查询出错")
	}
	serializer.Success(c, gin.H{"code": 200, "data": gameInfos}, "")
}

//新增游戏
func CreateGameInfo(c *gin.Context) {
	var createGameInfoDto dto.CreateGameInfoDto
	error := c.BindJSON(&createGameInfoDto)
	if error != nil {
		serializer.Response(c, http.StatusBadRequest, 500, nil, "系统异常")
	}
	gameInfo := &model.GameInfo{
		Id:         utils.IdWork(),
		GameName:   createGameInfoDto.GameName,
		GameCode:   createGameInfoDto.GameCode,
		ModelCode:  createGameInfoDto.ModelCode,
		GameStatus: createGameInfoDto.GameStatus,
		CreateTime: time.Now(),
		ModifyTime: time.Now(),
	}
	result := gameInfoService.CreateGameInfo(gameInfo)
	if result == "SUCCESS" {
		serializer.Success(c, gin.H{"code": 200, "data": result}, "")
	}
	serializer.Response(c, http.StatusBadRequest, 500, nil, "系统异常")
}

//游戏列表的基本信息
func UpdateGame(c *gin.Context) {
	var updateGameInfo dto.UpdateGameInfoDto
	error := c.BindJSON(&updateGameInfo)
	if error != nil {
		serializer.Response(c, http.StatusBadRequest, 500, nil, "系统异常")
	}
	result := gameInfoService.UpdateGameInfo(updateGameInfo.GameCode, updateGameInfo.GameStatus)
	if result == "SUCCESS" {
		serializer.Success(c, gin.H{"code": 200, "data": result}, "")
	}
	serializer.Fail(c, nil, "")
}

//获取游戏详情
func GameDetails(c *gin.Context) {
	gameCode := c.Query("gameCode")
	gameInfoDetailsList := gameBettingService.GameDetails(gameCode)
	if len(gameInfoDetailsList) == 0 {
		serializer.Fail(c, nil, "查询数据有误")
	}
	serializer.Success(c, gin.H{"code": 200, "data": gameInfoDetailsList}, "")
}

//投注项开关
func MathOddsFlag(c *gin.Context) {
	var bettingMathOddsFlgDto []*dto.BettingMathOddsFlgDto
	error := c.BindJSON(&bettingMathOddsFlgDto)
	if error != nil {
		serializer.Fail(c, nil, "")
	}
	result := gameBettingService.MathOddsFlag(bettingMathOddsFlgDto)
	if result == "SUCCESS" {
		serializer.Fail(c, nil, "")
	}
	serializer.Success(c, nil, "")
}
