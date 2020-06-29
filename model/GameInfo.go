package model

import (
	"github.com/shopspring/decimal"
	"reflect"
	"time"
)

//游戏信息实体
type GameInfo struct {
	Id int  `gorm:"PRIMARY_KEY"` //主键ID
	GameName string  //游戏名称
	ModelCode string //模型编码
	GameCode string  //游戏编码
	GameStatus string //游戏状态
	CreateTime time.Time //创建时间
	ModifyTime time.Time //修改时间
}

//投注项实体
type GameBetting struct {
	Id int `gorm:"PRIMARY_KEY"` //主键ID
	BettingName string
	BettingStatus string
	MathOdds decimal.Decimal
	CreateTime time.Time
	ModifyTime time.Time
	Attribute string
	LotteryAttribute string
}

//游戏分组实体
type GameGroup struct {
	Id int `gorm:"PRIMARY_KEY"` //主键ID
	GroupName string
	ModelCode string
	GroupStatus string
	CreateTime time.Time
	ModifyTime time.Time
}

//游戏模型实体
type GameModel struct {
	Id int `gorm:"PRIMARY_KEY"` //主键ID
	ModelName string
	ModelCode string
	ModelStatus string
	CreateTime time.Time
	ModifyTime time.Time
}

//游戏玩法实体
type GameItems struct {
	Id int `gorm:"PRIMARY_KEY"` //主键ID
	MethodName string
	ModelCode string
	GroupId int
	MethodStatus string
	CreateTime time.Time
	ModifyTime time.Time
}

//游戏计划实体
type GameScheduler struct {
	Id int `gorm:"PRIMARY_KEY"`
	GameName string `json:"gameName"`
	GameCode string `json:"gameCode"`
	ModelCode string `json:"modelCode"`
	DrawDay string `json:"drawDay"`
	DrawStime string `json:"drawStime"`
	DrawEtime string `json:"drawEtime"`
	OverAllTime int `json:"overallTime"`
	BetTime int `json:"betTime"`
	SealTime int `json:"sealTime"`
	CreateTime time.Time
	ModifyTime time.Time
}
func (gameScheduler GameScheduler) IsEmpty() bool {
	return reflect.DeepEqual(gameScheduler, GameScheduler{})
}

//开奖结果
type LotteryResults struct {
	Id int `gorm:"PRIMARY_KEY"`
	PeriodNum string
	GameCode string
	DrawTime string
	WinningResults string
	OutNumber string
	Status string
	IsClose string
	CreateTime time.Time
	ModifyTime time.Time
}
func (lotteryResults LotteryResults) IsEmpty() bool {
	return reflect.DeepEqual(lotteryResults, LotteryResults{})
}

type LotteryResultsDto struct {
	PeriodNum string
	StartTime string
	EndTime string
	CurrentPage int64
	PageSize int64
}


