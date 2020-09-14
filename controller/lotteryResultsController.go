package controller

import (
	"Game/common/dto"
	"Game/common/serializer"
	. "Game/service"
	"github.com/gin-gonic/gin"
	"time"
)

func GetLotteryResultList(c *gin.Context) {
	var lotteryResultsDto *dto.LotteryResultsDto
	err := c.BindJSON(&lotteryResultsDto)
	if err != nil {
		serializer.Fail(c, nil, err.Error())
	}
	var lotteryResultService = &LotteryResultService{}
	lotteryResultService.GetLotteryResultsList(lotteryResultsDto)
	serializer.Success(c, gin.H{"code": 200, "data": lotteryResultsDto}, "")
}

func CreatePeriodNum(ctx *gin.Context) {
	gameCode := ctx.Query("gameCode")
	modelCode := ctx.Query("modelCode")
	franchisee := ctx.Query("franchisee")
	var pk10 = &PK10Model{}
	pk10.CreatePeriodNum(gameCode, modelCode, franchisee)
}

func DrawOperator(ctx *gin.Context) {
	gameCode := ctx.Query("gameCode")
	modelCode := ctx.Query("modelCode")
	franchisee := ctx.Query("franchisee")
	var pk10 = &PK10Model{}
	time.Sleep(10 * time.Second)
	pk10.DrawOperator(gameCode, modelCode, franchisee)
}

func GetCurrentPeriod(ctx *gin.Context) {
	gameCode := ctx.Query("gameCode")
	modelCode := ctx.Query("modelCode")
	var pk10 = &PK10Model{}
	serializer.Success(ctx, gin.H{"data": pk10.GetCurrentPeriod(gameCode, modelCode)}, "")
}
