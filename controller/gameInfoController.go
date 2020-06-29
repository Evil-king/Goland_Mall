package controller

import (
	"Goland_Mall/dao"
	"Goland_Mall/dto"
	"Goland_Mall/model"
	"Goland_Mall/serializer"
	"Goland_Mall/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

//获取游戏列表
func PostGameInfoList(c *gin.Context) {
	var gameInfoDto dto.GameInfoDto
	error := c.BindJSON(&gameInfoDto)
	if error != nil {
		fmt.Printf("参数异常:%s", error.Error())
	}
	gameInfos, err := dao.GameInfoList(gameInfoDto)
	if err != nil {
		c.JSON(400, serializer.Fail(nil, err))
	}
	c.JSON(200, serializer.SuccessData(gameInfos))
}

//新增游戏
func CreateGameInfo(c *gin.Context) {
	var createGameInfoDto dto.CreateGameInfoDto
	error := c.BindJSON(&createGameInfoDto)
	if error != nil {
		fmt.Printf("参数异常:%s", error.Error())
		c.JSON(400, serializer.FailMsg("参数异常"))
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
	result := dao.CreateGameInfo(gameInfo)
	if result == "SUCCESS" {
		c.JSON(200, serializer.SuccessData(result))
	}
	c.JSON(400, serializer.FailMsg(""))
}

//游戏列表的基本信息
func UpdateGame(c *gin.Context) {
	var updateGameInfo dto.UpdateGameInfoDto
	error := c.BindJSON(&updateGameInfo)
	if error != nil {
		fmt.Printf("参数异常:%s", error.Error())
		c.JSON(400, serializer.FailMsg("参数异常"))
	}
	result := dao.UpdateGameInfo(updateGameInfo.GameCode, updateGameInfo.GameStatus)
	if result == "SUCCESS" {
		c.JSON(200, serializer.SuccessData(result))
	}
	c.JSON(400, serializer.FailMsg(""))
}

//获取游戏详情
func GameDetails(c *gin.Context) {
	gameCode := c.Query("gameCode")
	gameInfoDetailsList := dao.GameDetails(gameCode)
	if len(gameInfoDetailsList) == 0 {
		fmt.Println("查询数据有误")
		c.JSON(400, serializer.Fail("查询数据有误", nil))
	}
	c.JSON(200, serializer.SuccessData(gameInfoDetailsList))
}

//投注项开关
func MathOddsFlag(c *gin.Context) {
	var bettingMathOddsFlgDto []*dto.BettingMathOddsFlgDto
	error := c.BindJSON(&bettingMathOddsFlgDto)
	if error != nil {
		fmt.Printf("参数异常:%s", error.Error())
		c.JSON(400, serializer.Fail(nil, error))
	}
	result := dao.MathOddsFlag(bettingMathOddsFlgDto)
	if result == "SUCCESS" {
		c.JSON(400, serializer.Fail(result, nil))
	}
	c.JSON(200, serializer.Success())
}

