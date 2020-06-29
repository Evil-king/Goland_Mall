package main

import (
	"Goland_Mall/controller"
	"fmt"
	"github.com/gin-gonic/gin"
	"math/rand"
	"strconv"
	"time"
)

func main() {
	//conn := utils.RedisDefaultPool.Get()
	//result, err:=redis.String(conn.Do("get", "name"))
	//if err != nil{
	//	return
	//}
	//log.Println(result)

	//AA2006210001

	//str := time.Now().String()
	//str2 := strings.Split(str," ")[0]
	//str1 := strings.Replace(str2,"-","",-1)
	//fmt.Println(str1)
	//str3 :=str1[2:len(str1)]
	//fmt.Println(str3)

	//str := "AA2006210001"
	//newStr := str[8:len(str)]
	//fmt.Println(newStr)

	//var timeLayoutStr = "2006-01-02"
	//t:=time.Now()
	//str := t.Format(timeLayoutStr)
	//str1 :=strings.Replace(str,"-","",-1)[2:len(strings.Replace(str,"-","",-1))]
	//fmt.Println(str1)

	nums := generateRandomNumber(1, 11, 10)
	strSlice :=CalculationWiningResults(nums)
	fmt.Println(strSlice)
	str := Trans(strSlice)
	fmt.Println(str)
}

//生成count个[start,end)结束的不重复的随机数
func generateRandomNumber(start int, end int, count int) []int {
	//范围检查
	if end < start || (end-start) < count {
		return nil
	}
	//存放结果的slice
	nums := make([]int, 0)
	//随机数生成器，加入时间戳保证每次生成的随机数不一样
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for len(nums) < count {
		//生成随机数
		num := r.Intn((end - start)) + start
		//查重
		exist := false
		for _, v := range nums {
			if v == num {
				exist = true
				break
			}
		}
		if !exist {
			nums = append(nums, num)
		}
	}
	return nums
}

func CalculationWiningResults(nums []int) []string {
	lists := make([]string, 0)
	var one, two, three, four, five, six, seven, eight, nine, ten, sum int
	one = nums[0]
	two = nums[1]
	three = nums[2]
	four = nums[3]
	five = nums[4]
	six = nums[5]
	seven = nums[6]
	eight = nums[7]
	nine = nums[8]
	ten = nums[9]
	sum = one + ten
	lists = append(lists, strconv.Itoa(sum))
	if sum > 11 {
		lists = append(lists, "大")
	} else if sum <= 11 {
		lists = append(lists, "小")
	}
	if sum%2 == 0 {
		lists = append(lists, "双")
	} else {
		lists = append(lists, "单")
	}
	//1～5龙虎
	if one > ten {
		lists = append(lists, "龙")
	} else {
		lists = append(lists, "虎")
	}
	if two > nine {
		lists = append(lists, "龙")
	} else {
		lists = append(lists, "虎")
	}
	if three > eight {
		lists = append(lists, "龙")
	} else {
		lists = append(lists, "虎")
	}
	if four > seven {
		lists = append(lists, "龙")
	} else {
		lists = append(lists, "虎")
	}
	if five > six {
		lists = append(lists, "龙")
	} else {
		lists = append(lists, "虎")
	}
	return lists
}

func Trans(data interface{}) string {
	var str string
	if v,ok:=data.([]int);ok{
		str += "["
		for k, v := range v {
			if k == 0 {
				str = str + strconv.Itoa(v)
			} else {
				str = str + " " + strconv.Itoa(v)
			}
		}
		str += "]"
	} else if v,ok:=data.([]string);ok{
		str += "["
		for k, v := range v {
			if k == 0 {
				str = str + v
			} else {
				str = str + " " + v
			}
		}
		str += "]"
	}
	return str
}

func TransToString(winningResults []string) string {
	var str string
	str += "["
	for k, v := range winningResults {
		if k == 0 {
			str = str + v
		} else {
			str = str + " " + v
		}
	}
	str += "]"
	return str
}


func main1() {
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

	router.Run(":8000")
}
