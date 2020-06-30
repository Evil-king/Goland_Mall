package main

import (
	"Goland_Mall/controller"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/robfig/cron"
	"time"
)

func main1() {
	//conn := utils.RedisDefaultPool.Get()
	//result, err:=redis.String(conn.Do("get", "name"))
	//if err != nil{
	//	return
	//}
	//log.Println(result)

	//ch := make(chan int)
	//go recv(ch) // 启用goroutine从通道接收值
	//ch <- 10
	//fmt.Println("发送成功")

	//channel 练习
	//ch1 := make(chan int)
	//ch2 := make(chan int)
	//// 开启goroutine将0~100的数发送到ch1中
	//go func() {
	//	for i := 0; i <= 100; i++ {
	//		ch1 <- i
	//	}
	//	close(ch1)
	//}()
	//// 开启goroutine从ch1中接收值，并将该值的平方发送到ch2中
	//go func() {
	//	for { //开死循环一直不停的取 直到ch1没有值则跳出死循环
	//		i, ok := <-ch1
	//		if !ok {
	//			break
	//		}
	//		//并将该值的平方发送到ch2中
	//		ch2 <- i * i
	//	}
	//	close(ch2)
	//}()
	//// 在主goroutine中从ch2中接收值打印
	//for i := range ch2 { // 通道关闭后会退出for range循环
	//	fmt.Println(i)
	//}

	c := time.Unix(time.Now().UnixNano()/1e9, 0) //将秒转换为 time 类型
	fmt.Println(c)

}

//func recv(c chan int) {
//	ret := <-c
//	fmt.Println("接收成功", ret)
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
		v1.POST("/updateGameInfo", controller.UpdateGame)
		//游戏详情
		v1.GET("/gameDetails", controller.GameDetails)
		//投注项开关
		v1.POST("/updateMathOddsFlagBatch", controller.MathOddsFlag)
	}

	v2 := router.Group("/gameScheduler")
	{
		//获取游戏计划列表
		v2.GET("/list", controller.GameSchedulerList)
		//获取游戏计划列表
		v2.GET("/update", controller.GameSchedulerUpdate)
	}

	v3 := router.Group("/lotteryResults")
	{
		//获取游戏计划列表
		v3.POST("/list", controller.GetLotteryResultList)
	}
	// 定义一个cron运行器
	c := cron.New()
	// 定时5秒，每5秒执行print5
	c.AddFunc("0 0/1 * * * ?", controller.CreatePeriodNum)
	c.AddFunc("*/12 * * * * ?", controller.DrawOperator)
	c.Start()
	select {}

	router.Run(":8000")
}
