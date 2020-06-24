package controller

import (
	"Goland_Mall/dao"
	"Goland_Mall/dto"
	"Goland_Mall/model"
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
		fmt.Println(err)
		c.JSON(200, nil)
	}
	c.JSON(200, gameInfos)
}

//新增游戏
func CreateGameInfo(c *gin.Context) {
	var createGameInfoDto dto.CreateGameInfoDto
	error := c.BindJSON(&createGameInfoDto)
	if error != nil {
		fmt.Printf("参数异常:%s", error.Error())
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
		fmt.Println("result=", result)
		c.JSON(200, result)
	}
	c.JSON(400, nil)
}

//游戏列表的基本信息
func UpdateGame(c *gin.Context) {
	var updateGameInfo dto.UpdateGameInfoDto
	error := c.BindJSON(&updateGameInfo)
	if error != nil {
		fmt.Printf("参数异常:%s", error.Error())
	}
	result := dao.UpdateGameInfo(updateGameInfo.GameCode, updateGameInfo.GameStatus)
	if result == "SUCCESS" {
		fmt.Println("result=", result)
		c.JSON(200, result)
	}
	c.JSON(400, nil)
}

//获取游戏详情
func GameDetails(c *gin.Context) {
	gameCode := c.Query("gameCode")
	gameInfoDetailsList := dao.GameDetails(gameCode)
	if len(gameInfoDetailsList) == 0 {
		fmt.Println("查询数据有误")
	}
	c.JSON(200, gameInfoDetailsList)
}

func MathOddsFlag(c *gin.Context)  {
	var bettingMathOddsFlgDto []*model.BettingMathOddsFlgDto
	error := c.BindJSON(&bettingMathOddsFlgDto)
	if error != nil {
		fmt.Printf("参数异常:%s", error.Error())
	}
	result := dao.MathOddsFlag(bettingMathOddsFlgDto)
	if result == "SUCCESS" {
		fmt.Println("result=", result)
		c.JSON(200, result)
	}
	c.JSON(400, nil)
}