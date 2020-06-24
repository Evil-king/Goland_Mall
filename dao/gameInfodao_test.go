package dao

import (
	"Goland_Mall/model"
	"fmt"
	"testing"
	"time"
)

func TestGameInfo(t *testing.T)  {
	fmt.Println("测试gameInfo中的相关函数")
}


func TestCreateGameInfo(t *testing.T) {
	gameInfo := &model.GameInfo{
		Id:1255444774531198909,
		GameName: "东京热",
		GameCode: "BB",
		ModelCode: "PK10",
		GameStatus: "1",
		CreateTime: time.Now(),
		ModifyTime: time.Now(),
	}
	CreateGameInfo(gameInfo)
}
