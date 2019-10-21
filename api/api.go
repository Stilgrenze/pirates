package api

import "Pirates/game"

type Api interface {
	Init(game *game.Game)
}

func NewApi() Api {
	api := new(HttpApi)
	return api
}
