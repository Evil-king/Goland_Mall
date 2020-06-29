package controller

import (
	"Goland_Mall/dto"
	"Goland_Mall/serializer"
	"Goland_Mall/service"
	"github.com/gin-gonic/gin"
)

//开奖计划列表
func GameSchedulerList(c *gin.Context)  {
	result := service.GameSchedulerList()
	c.JSON(200, result)
}

//更新开奖计划
func GameSchedulerUpdate(c *gin.Context)  {
	var gameSchedulerDto dto.GameSchedulerDto
	error := c.BindJSON(&gameSchedulerDto)
	if error != nil {
		c.JSON(400, serializer.Fail(nil, error))
	}
	result := service.GameSchedulerUpdate(gameSchedulerDto)
	if result.Msg == "success" {
		c.JSON(200, serializer.SuccessData(result))
	}
	c.JSON(400, serializer.SuccessData(result))
}