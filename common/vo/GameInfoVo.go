package vo

import (
	"Game/model"
	"fmt"
	"github.com/shopspring/decimal"
	"time"
)

type GameInfoDetailsVo struct {
	GroupName     string          `json:"groupName"`
	MethodName    string          `json:"methodName"`
	MathOdds      decimal.Decimal `json:"mathOdds"`
	BettingName   string          `json:"bettingName"`
	Id            int             `json:"bettingId"`
	BettingStatus string          `json:"bettingStatus"`
}

type GameInfoInnerVO struct {
	GameCode string
	GameName string
	GroupName string
	GameMode string
	MathOdds string
	BettingName string
	BettingStatus string
	Id string
	MethodName string
	Attributes string
	LotteryAttributes string
	PageAttributes string
	Sort string
}

type LotteryResultsListVo struct {
	PeriodNum      string
	DrawTime       time.Time
	WinningResults string
	OutNumber      string
}

type GameInfoInnerAggregation struct {
	GameName          string          `json:"gameName"`
	ModelCode         string          `json:"modelCode"`
	GameStatus        string          `json:"gameStatus"`
	GroupName         string          `json:"groupName"`
	MathOdds          decimal.Decimal `json:"mathOdds"`
	BettingName       string          `json:"bettingName"`
	BettingStatus     string          `json:"bettingStatus"`
	Id         		  string          `json:"bettingId"`
	MethodName        string          `json:"methodName"`
	Attributes        string          `json:"attributes"`
	LotteryAttributes string          `json:"lotteryAttributes"`
}

type NewResultVo struct {
	WiningResult string `json:"winingResult"`
	NextPeriodNum string `json:"nextPeriodNum"`
	SealTime string `json:"sealTime"`
	OutNumber string `json:"outNumber"`
	GameCode string `json:"gameCode"`
	EndTime int64 `json:"endTime"`
	IsClose string `json:"isClose"`
}

type CurrentPeriodVo struct {
	PeriodNum string `json:"periodNum"`
	EndTime int64 `json:"endTime"`
	SealTime string `json:"sealTime"`
	IsClose string `json:"isClose"`
}

type GameSchedulerVO struct {
	Id int64 `gorm:"PRIMARY_KEY"`
	GameName string `json:"gameName"`
	GameCode string `json:"gameCode"`
	ModelCode string `json:"modelCode"`
	DrawDay string `json:"drawDay"`
	DrawStime string `json:"drawStime"`
	DrawEtime string `json:"drawEtime"`
	OverAllTime int `json:"overallTime"`
	BetTime int `json:"betTime"`
	SealTime int `json:"sealTime"`
}

func (s GameSchedulerVO)GameSchedulerToVo(gameScheduler model.GameScheduler) *GameSchedulerVO {
	return &GameSchedulerVO{
		Id:          gameScheduler.Id,
		GameName:    gameScheduler.GameName,
		GameCode:    gameScheduler.GameCode,
		ModelCode:   gameScheduler.ModelCode,
		DrawDay:     gameScheduler.DrawDay,
		DrawStime:   gameScheduler.DrawStime,
		DrawEtime:   gameScheduler.DrawEtime,
		OverAllTime: gameScheduler.OverAllTime,
		BetTime:     gameScheduler.BetTime,
		SealTime:    gameScheduler.SealTime,
	}
}

type jsonTime time.Time

//实现它的json序列化方法
func (this jsonTime) MarshalJSON() ([]byte, error) {
	var stamp = fmt.Sprintf("\"%s\"", time.Time(this).Format("2006-01-02 15:04:05"))
	return []byte(stamp), nil
}