package dto

type GameSchedulerDto struct {
	GameCode string `json:"gameCode"`
	DrawStartTime string `json:"drawStartTime"`
	DrawEndTime string  `json:"drawEndTime"`
	OverallTime int `json:"overallTime"`
	SealTime int `json:"sealTime"`
}