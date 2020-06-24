package controller

import (
	"Goland_Mall/dao"
	"Goland_Mall/dto"
	"Goland_Mall/model"
	"Goland_Mall/serializer"
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
	gameInfos, err := dao.PostGameInfoList(gameInfoDto)
	if err != nil {
		c.JSON(400, serializer.Fail(nil, err))
	}
	c.JSON(200, serializer.Success(gameInfos, nil))
}

//新增游戏
func CreateGameInfo(c *gin.Context) {
	var createGameInfoDto dto.CreateGameInfoDto
	error := c.BindJSON(&createGameInfoDto)
	if error != nil {
		fmt.Printf("参数异常:%s", error.Error())
		c.JSON(400, serializer.Success(nil, error))
	}
	gameInfo := &model.GameInfo{
		//TODO 这里id要换成自动生成
		Id:         1255444774531198909,
		GameName:   createGameInfoDto.GameName,
		GameCode:   createGameInfoDto.GameCode,
		ModelCode:  createGameInfoDto.ModelCode,
		GameStatus: createGameInfoDto.GameStatus,
		CreateTime: time.Now(),
		ModifyTime: time.Now(),
	}
	result := dao.CreateGameInfo(gameInfo)
	if result == "SUCCESS" {
		c.JSON(200, serializer.Success(result, nil))
	}
	c.JSON(400, serializer.Success(result, nil))
}

//游戏列表的基本信息
func UpdateGame(c *gin.Context) {
	var updateGameInfo dto.UpdateGameInfoDto
	error := c.BindJSON(&updateGameInfo)
	if error != nil {
		fmt.Printf("参数异常:%s", error.Error())
		c.JSON(400, serializer.Success(nil, error))
	}
	result := dao.UpdateGameInfo(updateGameInfo.GameCode, updateGameInfo.GameStatus)
	if result == "SUCCESS" {
		c.JSON(200, serializer.Success(result, nil))
	}
	c.JSON(400, serializer.Success(result, nil))
}

//获取游戏详情
func GameDetails(c *gin.Context) {
	gameCode := c.Query("gameCode")
	gameInfoDetailsList := dao.GameDetails(gameCode)
	if len(gameInfoDetailsList) == 0 {
		fmt.Println("查询数据有误")
		c.JSON(400, serializer.Fail("查询数据有误", nil))
	}
	c.JSON(200, serializer.Success(gameInfoDetailsList, nil))
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
	c.JSON(200, serializer.Success(nil, nil))
}

//开奖计划列表
func GameSchedulerList(c *gin.Context)  {
	gameSchedulerList := dao.GameSchedulerList()
	if gameSchedulerList == nil{
		c.JSON(400, serializer.Fail(gameSchedulerList, nil))
	}
	c.JSON(200, serializer.Success(gameSchedulerList, nil))
}

//更新开奖计划
func GameSchedulerUpdate(c *gin.Context)  {
	var gameSchedulerDto dto.GameSchedulerDto
	error := c.BindJSON(&gameSchedulerDto)
	if error != nil {
		c.JSON(400, serializer.Fail(nil, error))
	}
	result := dao.GameSchedulerUpdate(gameSchedulerDto)
	if result == "SUCCESS" {
		c.JSON(200, serializer.Success(result, nil))
	}
	c.JSON(400, serializer.Success(result, nil))
}