package model

import "github.com/shopspring/decimal"

type ZodiacOddsRelation struct {
	Id int64 `gorm:"PRIMARY_KEY"`
	ItemsName string `json:"itemsName"`
	BettingName string `json:"bettingName"`
	NatalOdds decimal.Decimal `json:"natalOdds"`
	NoNatalOdds decimal.Decimal `json:"noNatalOdds"`
}
