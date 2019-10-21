package game

import (
	"Pirates/graphic"
	"Pirates/models/info"
	"Pirates/models/ocean"
	"Pirates/models/player"
	"time"
)

type Game struct {
	tick           int
	renderer       graphic.Renderer
	timer          Timer
	updateListener []func()

	// Game Models
	ocean   *ocean.Ocean
	players map[string]*player.Player
}

func NewGame(renderer graphic.Renderer, ocean *ocean.Ocean, tick int) *Game {
	g := new(Game)
	g.updateListener = make([]func(), 0, 0)
	g.renderer = renderer
	g.ocean = ocean
	g.tick = tick
	g.players = make(map[string]*player.Player)
	return g
}

func (g *Game) GetOcean() ocean.Ocean {
	return *g.ocean
}

func (g *Game) GetPlayers() []info.InfoPlayer {
	players := make([]info.InfoPlayer, 0, 0)
	for _, player := range g.players {
		players = append(players, info.InfoPlayer{
			player.Name,
			player.Gold,
			player.GoldSpent,
		})
	}
	return players
}

func (g *Game) RegisterUpdateListener(listener func()) {
	g.updateListener = append(g.updateListener, listener)
}

func (g *Game) MainLoop() {
	g.timer.Start()

	for {
		g.timer.Update()
		g.updateListeners()
		g.ocean.Tick()
		g.goldToPlayers()
		if g.renderer != nil {
			g.renderer.Draw(g.ocean)
		}
		time.Sleep(time.Duration(g.tick) * time.Millisecond)
	}
}

func (g *Game) goldToPlayers() {
	for _, player := range g.players {
		for _, ship := range player.Ships {
			if ship.Gold > 0 {
				player.Gold += int64(ship.Gold)
				ship.Gold = 0
			}
		}
	}
}

func (g *Game) updateListeners() {
	for i := 0; i < len(g.updateListener); i++ {
		g.updateListener[i]()
	}
}
