package controller

import (
	"Game/common/dto"
	"Game/common/serializer"
	. "Game/service"
	"github.com/gin-gonic/gin"
)

var gameSchedulerService  = &GameSchedulerService{}

//开奖计划列表
func GameSchedulerList(c *gin.Context) {
	result := gameSchedulerService.GameSchedulerList
	serializer.Success(c, gin.H{"code": 200, "data": result}, "")
}

//更新开奖计划
func GameSchedulerUpdate(c *gin.Context) {
	var gameSchedulerDto dto.GameSchedulerDto
	error := c.BindJSON(&gameSchedulerDto)
	if error != nil {
		serializer.Fail(c, nil, error.Error())
	}
	flag := gameSchedulerService.GameSchedulerUpdate(gameSchedulerDto)
	if flag {
		serializer.Success(c, nil, "更新成功")
	}
	serializer.Fail(c, nil, "更新失败")
}
