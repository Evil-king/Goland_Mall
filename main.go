package main

import (
	"Goland_Mall/controller"
	"github.com/gin-gonic/gin"
)

//func main() {
//	conn := utils.RedisDefaultPool.Get()
//	result, err:=redis.String(conn.Do("get", "name"))
//	if err != nil{
//		return
//	}
//	log.Println(result)
//}

func main() {
	router := gin.Default()

	router.Use(gin.Recovery())

	v1 := router.Group("/game")
	{
		//获取游戏列表方法
		v1.POST("/gameInfo", controller.PostGameInfoList)
		//创建游戏方法
		v1.POST("/createGameInfo", controller.CreateGameInfo)
		//更改游戏状态
		v1.POST("/updateGameInfo",controller.UpdateGame)
		//游戏详情
		v1.GET("/gameDetails",controller.GameDetails)
		//投注项开关
		v1.POST("/updateMathOddsFlagBatch",controller.MathOddsFlag)
	}

	router.Run(":8000")
}
