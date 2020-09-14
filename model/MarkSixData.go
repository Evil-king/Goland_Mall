package model

import "reflect"

type MarkSixData struct {
	Id int64 `gorm:"PRIMARY_KEY"`
	Year string
	Month string
	Day string
	Data string
	ChineseZodiac string
}

func (markSixData MarkSixData) IsEmpty() bool {
	return reflect.DeepEqual(markSixData, MarkSixData{})
}