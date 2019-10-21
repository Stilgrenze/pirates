package main

import (
	"Pirates/api"
	"Pirates/events"
	game2 "Pirates/game"
	_ "Pirates/graphic/text"
	"Pirates/models/ocean"
	"math/rand"
	"time"
)

var game *game2.Game

func main() {
	rand.Seed(time.Now().UnixNano())

	initGame()
	initApi()
}

func initGame() {
	events.Init()
	events.GetInstance().CreateEvent("Game started!", 0, 0)
	ocean := ocean.GenerateOcean(20, 16)
	//renderer := text.TextRenderer{}
	game = game2.NewGame(nil, ocean, 1000)
	go game.MainLoop()
}

func initApi() {
	api := api.NewApi()
	api.Init(game)
}
