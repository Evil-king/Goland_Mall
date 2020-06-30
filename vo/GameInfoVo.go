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
	DrawTime       *time.Time
	WinningResults string
	OutNumber      string
}
