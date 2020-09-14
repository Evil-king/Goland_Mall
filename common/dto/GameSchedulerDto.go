package dto

type GameSchedulerDto struct {
	GameCode string `json:"gameCode"`
	ModelCode string `json:"modelCode"`
	DrawStartTime string `json:"drawStartTime"`
	DrawEndTime string  `json:"drawEndTime"`
	OverallTime int `json:"overallTime"`
	SealTime int `json:"sealTime"`
	BetTime int `json:"betTime"`
	DrawDay []*string `json:"drawDay"`
	DrawTime string `json:"drawTime"`
	DrawCycle string `json:"drawCycle"`
}