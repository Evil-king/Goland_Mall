package dto

type GameInfoDto struct {
	GameName string
	GameStatus string
	ModelCode string
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