package model

import (
	"math/big"
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
	MathOdds big.Rat
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

type BettingMathOddsFlgDto struct {
	BettingId string
	Flag string
}
