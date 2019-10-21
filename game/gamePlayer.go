package game

import (
	"Pirates/api/request"
	"Pirates/models/player"
)

func (g *Game) AddPlayer(name, secret string) *player.Player {
	newPlayer := player.Player{
		Name:   name,
		Secret: secret,
		Gold:   5000,
	}

	if _, ok := g.players[name]; ok == false {
		g.players[name] = &newPlayer
	} else {
		return nil
	}

	return &newPlayer
}

func (g *Game) GetPlayerWithRequest(request request.PlayerRequest) *player.Player {
	if _, ok := g.players[request.Name]; ok {
		if g.players[request.Name].Secret == request.Secret {
			return g.players[request.Name]
		}
	}
	return nil
}
