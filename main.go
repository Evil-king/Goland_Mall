package main

import (
	"Game/controller"
	"Game/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"log"
	"os"
	"os/signal"
	"reflect"
	"syscall"
)

func main1() {
	//utils.SendMsg("Topic-Test","First to message")

	//res, err := redis.Int64(conn.Do("HGET", "student","age"))
	//if err != nil {
	//	fmt.Println("redis HGET error:", err)
	//} else {
	//	fmt.Printf("res  : %d \n", res)
	//}

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
}
//func recv(c chan int) {
//	ret := <-c
//	fmt.Println("接收成功", ret)
//}

func Duplicate(a interface{}) (ret []interface{}) {
	va := reflect.ValueOf(a)
	for i := 0; i < va.Len(); i++ {
		if i > 0 && reflect.DeepEqual(va.Index(i-1).Interface(), va.Index(i).Interface()) {
			continue
		}
		ret = append(ret, va.Index(i).Interface())
	}
	return ret
}

func removeDuplicateElement(addrs []string) []string {
	result := make([]string, 0, len(addrs))
	temp := map[string]struct{}{}
	for _, item := range addrs {
		if _, ok := temp[item]; !ok {
			temp[item] = struct{}{}
			result = append(result, item)
		}
	}
	return result
}

// @title Swagger Example API
// @version 1.0
// @description This is a sample server celler server.
// @termsOfService https://razeen.me

// @contact.name Razeen
// @contact.url https://razeen.me
// @contact.email me@razeen.me

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host 127.0.0.1:8080
// @BasePath /api/v1
func main() {
	GrpcInit()
	router := gin.Default()
	router.Use(gin.Recovery())
	//swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	//健康检查
	router.GET("/health", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"status": "OK",
		})
	})

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
		v3.GET("/getCurrentPeriod", controller.GetCurrentPeriod)
	}

	v4 := router.Group("/timer")
	{
		v4.GET("/createPeriodNum",controller.CreatePeriodNum)
		v4.GET("/drawOperator",controller.DrawOperator)
	}

	errChan := make(chan error)
	//启动携程运行启动
	go func() {
		utils.RegService() //注册服务
		err :=router.Run(":8000")
		if err != nil {
			log.Panicln(err)
			errChan <- err
		}
	}()
	//服务优雅的关闭
	go func() {
		sig_c := make(chan os.Signal)
		signal.Notify(sig_c, syscall.SIGINT, syscall.SIGTERM)
		errChan <- fmt.Errorf("%s", <-sig_c)
	}()
	getErr := <-errChan
	//关闭注册
	utils.Unregservice()
	log.Println(getErr)
}
