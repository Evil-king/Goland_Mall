package vo

import "github.com/shopspring/decimal"

type GameInfoDetailsVo struct {
	GroupName string `json:"groupName"`
	MethodName string `json:"methodName"`
	MathOdds decimal.Decimal `json:"mathOdds"`
	BettingName string 	`json:"bettingName"`
	Id int `json:"bettingId"`
	BettingStatus string `json:"bettingStatus"`
}
