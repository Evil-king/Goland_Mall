package service

import (
	. "Game/common/vo"
	"Game/dao"
	"Game/model"
	"Game/utils"
	"bytes"
	"encoding/json"
	"github.com/garyburd/redigo/redis"
	"log"
	"strconv"
	"time"
)

var gameSchedulerService = &GameSchedulerService{}
var gameUtil = &GameUtils{}

type IPK10Model interface {
	DrawOperator(gameCode string) string
	GetNewLotteryResults(gameCode string) NewResultVo
	CreatePeriodNum(gameCode string) string
	GetCurrentPeriod(gameCode string, modelCode string) CurrentPeriodVo
	//UpdateSchedulerByTime()
	InsertData(gameCode string, periodNum string, outNumber []int, winningResults []string, flag string)
	updateLotteryResult(lotteryResults model.LotteryResults, scheduler model.GameScheduler, outNumber []int, winningResults []string, status string, flag string)
}
type PK10Model struct{}

//更新期号对应的开出号码和开奖结果
func (p *PK10Model) DrawOperator(gameCode string, modelCode string, franchisee string) string {
	log.Println("进入DrawOperator方法", gameCode)
	var timeLayoutStr = "15:04"
	ts := time.Now().Format(timeLayoutStr) //time转string
	//通过gameCode获取游戏计划
	gameScheduler := gameSchedulerService.GetSchedulerByGameCode(gameCode)
	//判断当前正在开的期号是否已经开奖
	lotteryResults := dao.SelectPeriodNumByStatus(gameCode)
	if lotteryResults.IsEmpty() {
		return "没有开奖期号"
	}
	log.Println("lotteryResults", lotteryResults)
	flag := gameUtil.IsEffectiveDateStr(ts, gameScheduler.DrawStime, gameScheduler.DrawEtime) == false
	if "everyday" == gameScheduler.DrawDay && flag {
		if "true" == lotteryResults.IsClose {
			//step1、随机开出十个数
			outNumber := gameUtil.RandomNumber(modelCode)
			//step2、根据不同规则算出最终的开奖结果
			winningResults := gameUtil.CalculationWiningResults(modelCode, outNumber)
			result := gameUtil.OperatorWingData(gameCode, modelCode, outNumber)
			//step3、更新表
			p.updateLotteryResult(lotteryResults, gameScheduler, outNumber, winningResults, "close", "false")
			//TODO step4、组装开奖结果、发送mq消息给注单
			mp := make(map[string]interface{})
			mp[lotteryResults.PeriodNum+"#"+franchisee+"#"+gameScheduler.ModelCode] = result
			byteBuf := bytes.NewBuffer([]byte{})
			encoder := json.NewEncoder(byteBuf)
			encoder.SetEscapeHTML(false)
			err := encoder.Encode(mp)
			if err != nil {
				log.Fatal(err)
			}
			utils.SendMsg("Topic-game", byteBuf.String())
		} else {
			return "当前时间不在游戏开放时间之内，不允许开奖"
		}
	} else {
		//step1、随机开出十个数
		outNumber := gameUtil.RandomNumber(modelCode)
		//step2、根据不同规则算出最终的开奖结果
		winningResults := gameUtil.CalculationWiningResults(modelCode, outNumber)
		//组装发送mq消息内容
		result := gameUtil.OperatorWingData(gameCode, modelCode, outNumber)
		//step3、更新表
		p.updateLotteryResult(lotteryResults, gameScheduler, outNumber, winningResults, "open", "true")
		//TODO step4、组装开奖结果、发送mq消息给注单
		mp := make(map[string]interface{})
		mp[lotteryResults.PeriodNum+"#"+franchisee+"#"+gameScheduler.ModelCode] = result
		byteBuf := bytes.NewBuffer([]byte{})
		encoder := json.NewEncoder(byteBuf)
		encoder.SetEscapeHTML(false)
		err := encoder.Encode(mp)
		if err != nil {
			log.Fatal(err)
		}
		utils.SendMsg("Topic-game", byteBuf.String())
	}
	return "SUCCESS"
}

//创建期号并且入库
func (p *PK10Model) CreatePeriodNum(gameCode string, modelCode string, franchisee string) string {
	log.Println("进入CreatePeriodNum方法", gameCode)
	//通过gameCode获取游戏计划
	gameScheduler := gameSchedulerService.GetSchedulerByGameCode(gameCode)
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
		p.InsertData(gameCode, modelCode, periodNum, nil, nil, "true", franchisee)
		return "SUCCESS"
	}
	newPeriodNum := utils.AddPeriodNum(lotteryResults.PeriodNum, gameCode)
	//写入表中
	p.InsertData(gameCode, modelCode, newPeriodNum, nil, nil, "true", franchisee)
	return "SUCCESS"
}

func (p *PK10Model) InsertData(gameCode string, modelCode string, periodNum string, outNumber []int,
	winningResults []string, flag string, franchisee string) {
	dao.InsertData(gameCode, modelCode, periodNum, outNumber, winningResults, flag)
	//发送mq告诉注单 已经封盘
	message := franchisee + "#" + periodNum
	utils.SendMsg("Topic-game", message)
}

//更新期号对应的开奖号码和结果
func (p *PK10Model) updateLotteryResult(lotteryResults model.LotteryResults, scheduler model.GameScheduler,
	outNumber []string, winningResults []string, status string, flag string) {
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
	x := time.Unix(lastIssueDrawTime/1000, 1)
	log.Println("time={}", x)
	newLotteryResult := model.LotteryResults{
		Id:             lotteryResults.Id,
		DrawTime:       x,
		WinningResults: utils.SliceToString(winningResults),
		OutNumber:      utils.SliceToString(outNumber),
		IsClose:        flag,
		Status:         status,
		ModifyTime:     time.Now(),
	}
	dao.UpdateLotteryByParams(newLotteryResult)
}

func (p *PK10Model) GetNewLotteryResults(gameCode string, modelCode string) NewResultVo {
	var gameSchedulerService = &GameSchedulerService{}
	var lotteryResultService = &LotteryResultService{}
	//获取游戏计划
	gameScheduler := gameSchedulerService.SelectGameSchedulerToEveryDay(gameCode, modelCode)
	//获取当前期号
	lotteryResults := lotteryResultService.GetCurrentPeriod(gameCode)
	//获取下起期号
	nextPeriodNum := utils.AddPeriodNum(lotteryResults.PeriodNum, gameCode)

	overallTime := int64((gameScheduler.SealTime + gameScheduler.BetTime) * 1000)
	createTime := lotteryResults.DrawTime.UnixNano() / 1e6
	return NewResultVo{
		WiningResult:  lotteryResults.WinningResults,
		NextPeriodNum: nextPeriodNum,
		SealTime:      string(gameScheduler.SealTime),
		OutNumber:     lotteryResults.OutNumber,
		GameCode:      gameCode,
		IsClose:       lotteryResults.IsClose,
		EndTime:       overallTime + createTime,
	}
}

func (p *PK10Model) GetCurrentPeriod(gameCode string, modelCode string) CurrentPeriodVo {
	var gameSchedulerService = &GameSchedulerService{}
	var lotteryResultService = &LotteryResultService{}
	//获取游戏计划
	gameScheduler := gameSchedulerService.SelectGameSchedulerToEveryDay(gameCode, modelCode)
	//获取当前期号
	lotteryResults := lotteryResultService.GetCurrentPeriod(gameCode)
	var isClose = ""
	//获取下起期号
	nextPeriodNum := utils.AddPeriodNum(lotteryResults.PeriodNum, gameCode)
	sumNum := gameScheduler.CreateTime.UnixNano()/1e6 + int64(gameScheduler.OverAllTime)
	nowTime := time.Now().UnixNano() / 1e6
	if sumNum < nowTime {
		isClose = "false"
	} else {
		isClose = lotteryResults.IsClose
	}
	//封盘时长 先从缓存中读取
	var overallTime int64
	conn := utils.RedisDefaultPool.Get()
	if _, err := redis.Bool(conn.Do("EXISTS", gameCode+modelCode)); err != nil {
		sealTime, err := redis.Int64(conn.Do("HGET", gameCode+modelCode, "sealTime"))
		betTime, err := redis.Int64(conn.Do("HGET", gameCode+modelCode, "betTime"))
		if err != nil {
			log.Fatal(err)
		}
		overallTime = sealTime + betTime
	} else {
		overallTime = int64(gameScheduler.OverAllTime)
	}
	var endTime int64
	if !lotteryResults.IsEmpty() {
		if "true" == lotteryResults.IsClose &&
			(("open" == lotteryResults.Status) || "close" == lotteryResults.Status) {
			endTime = lotteryResults.DrawTime.UnixNano()/1e6 + overallTime
		} else {
			return CurrentPeriodVo{
				PeriodNum: "",
				EndTime:   0,
				SealTime:  strconv.Itoa(gameScheduler.SealTime),
				IsClose:   "false",
			}
		}
	} else {
		return CurrentPeriodVo{
			PeriodNum: "",
			EndTime:   0,
			SealTime:  strconv.Itoa(gameScheduler.SealTime),
			IsClose:   "false",
		}
	}
	return CurrentPeriodVo{
		PeriodNum: nextPeriodNum,
		EndTime:   endTime,
		SealTime:  strconv.Itoa(gameScheduler.SealTime),
		IsClose:   isClose,
	}
}
