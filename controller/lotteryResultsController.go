package controller

import (
	"Goland_Mall/model"
	"Goland_Mall/serializer"
	"Goland_Mall/service"
	"github.com/gin-gonic/gin"
)

func GetLotteryResultList(c *gin.Context)  {
	var lotteryResultsDto *model.LotteryResultsDto
	err :=c.BindJSON(&lotteryResultsDto)

	if err !=nil{
		c.JSON(400,serializer.Fail(nil,err))
	}

	c.JSON(200,service.GetLotteryResultsList(lotteryResultsDto))
}
