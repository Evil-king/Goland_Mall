package vo

import (
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
	BettingId         string          `json:"bettingId"`
	MethodName        string          `json:"methodName"`
	Attributes        string          `json:"attributes"`
	LotteryAttributes string          `json:"lotteryAttributes"`
}
