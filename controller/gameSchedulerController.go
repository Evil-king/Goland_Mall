package controller

import (
	"Goland_Mall/dao"
	"Goland_Mall/dto"
	"Goland_Mall/serializer"
	"github.com/gin-gonic/gin"
)

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