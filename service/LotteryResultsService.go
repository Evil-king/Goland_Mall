package service

import (
	"Goland_Mall/dao"
	"Goland_Mall/model"
	"Goland_Mall/serializer"
	"Goland_Mall/utils"
	"log"
	"math/rand"
	"strconv"
	"time"
)

//获取开奖结果
func GetLotteryResultsList(dto *model.LotteryResultsDto) serializer.Result {
	result := dao.GetLotteryResultsList(dto)
	return serializer.SuccessData(result)
}

func DrawOperator(gameCode string) serializer.Result {
	log.Println("进入DrawOperator方法", gameCode)
	var timeLayoutStr = "15:04"
	ts := time.Now().Format(timeLayoutStr) //time转string
	//通过gameCode获取游戏计划
	gameScheduler := GetSchedulerByGameCode(gameCode)
	//判断当前正在开的期号是否已经开奖
	lotteryResults := dao.SelectPeriodNumByStatus(gameCode)
	log.Println("lotteryResults", lotteryResults)
	if "everyday" == gameScheduler.DrawDay && IsEffectiveDateStr(ts, gameScheduler.DrawStime, gameScheduler.DrawEtime) == false {
		if "true" == lotteryResults.IsClose {
			//step1、随机开出十个数
			outNumber := generateRandomNumber(1, 11, 10)
			//step2、根据不同规则算出最终的开奖结果
			winningResults := CalculationWiningResults(outNumber)
			//step3、更新表
			updateLotteryResult(lotteryResults, gameScheduler, outNumber, winningResults, "false", "close")
			//step4、组装开奖结果、发送mq消息给注单
		} else {
			return serializer.FailMsg("当前时间不在游戏开放时间之内，不允许开奖")
		}
	} else {
		//step1、随机开出十个数
		outNumber := generateRandomNumber(1, 11, 10)
		//step2、根据不同规则算出最终的开奖结果
		winningResults := CalculationWiningResults(outNumber)
		//step3、更新表
		updateLotteryResult(lotteryResults, gameScheduler, outNumber, winningResults, "true", "open")
		//step4、组装开奖结果、发送mq消息给注单
	}
	return serializer.Success()
}

func CreatePeriodNum(gameCode string) serializer.Result {
	log.Println("进入CreatePeriodNum方法", gameCode)
	//通过gameCode获取游戏计划
	gameScheduler := GetSchedulerByGameCode(gameCode)
	if gameScheduler.IsEmpty() {
		panic("该游戏没有开奖计划")
	}
	lotteryResults := dao.SelectPeriodNumByIsClose(gameCode)
	var periodNum string
	if lotteryResults.IsEmpty() {
		log.Println("进入次方法")
		//期数规则：gameCode+200503+0001
		periodNum = utils.GetPeriodNum("true", periodNum, gameCode)
		//写入表中
		InsertData(gameCode, periodNum, nil, nil, "true")
		return serializer.Success()
	}
	newPeriodNum := utils.AddPeriodNum(lotteryResults.PeriodNum, gameCode)
	//写入表中
	InsertData(gameCode, newPeriodNum, nil, nil, "true")
	return serializer.Success()
}

func InsertData(gameCode string, periodNum string, outNumber []int, winningResults []string, flag string) {
	dao.InsertData(gameCode, periodNum, outNumber, winningResults, flag)
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

func IsEffectiveDateStr(nowTime string, sTime string, eTime string) bool {
	var timeLayoutStr = "15:04"
	st, _ := time.ParseInLocation(timeLayoutStr, sTime, time.Local)   //string转time
	et, _ := time.ParseInLocation(timeLayoutStr, eTime, time.Local)   //string转time
	nt, _ := time.ParseInLocation(timeLayoutStr, nowTime, time.Local) //string转time
	log.Println("判断时间是否在计划时间内", st.Before(nt) && et.After(nt))
	return st.Before(nt) && et.After(nt)
}

func updateLotteryResult(lotteryResults model.LotteryResults, scheduler model.GameScheduler, outNumber []int, winningResults []string, status string, flag string) {
	//获取上期的期号
	lastPeriodNum := utils.LastIssuePeriodNum(lotteryResults.PeriodNum, lotteryResults.GameCode)
	log.Println("lastIssuePeriodNum={}", lastPeriodNum)
	//获取上一期开奖内容
	lastIssueLotteryResults := dao.SelectLastIssue(lotteryResults.PeriodNum)
	var lastIssueDrawTime int64
	if lastIssueLotteryResults.IsEmpty() {
		lastIssueDrawTime = time.Now().UnixNano() / 1e6
	} else if lastIssueLotteryResults.CreateTime.UnixNano()/1e6+int64(scheduler.OverAllTime*1000) < time.Now().UnixNano()/1e6 {
		lastIssueDrawTime = time.Now().UnixNano() / 1e6
	} else {
		lastIssueDrawTime = lastIssueLotteryResults.CreateTime.UnixNano()/1e6 + int64(scheduler.OverAllTime*1000)
	}
	log.Println("lastIssueDrawTime={}", lastIssueDrawTime)
	value, _ := time.ParseInLocation("2006-01-02 15:04:05", string(lastIssueDrawTime), time.Local)
	log.Println("time={}", value)
	newLotteryResult := model.LotteryResults{
		Id:             lotteryResults.Id,
		DrawTime:       value,
		WinningResults: utils.SliceToString(winningResults),
		OutNumber:      utils.SliceToString(outNumber),
		IsClose:        flag,
		Status:         status,
		ModifyTime:     time.Now(),
	}
	dao.UpdateLotteryByParams(newLotteryResult)
}
