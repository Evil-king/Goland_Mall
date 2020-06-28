package dto

type GameInfoDto struct {
	GameName string
	GameStatus string
	ModelCode string
	CurrentPage int64 //当前页
	PageSize int64 //每页记录数
}

type CreateGameInfoDto struct {
	GameName string
	GameCode string
	ModelCode string
	GameStatus string
}

type UpdateGameInfoDto struct {
	GameCode string
	GameStatus string
}

type BettingMathOddsFlgDto struct {
	BettingId string
	Flag string
}