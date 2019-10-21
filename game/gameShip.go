package game

import (
	"Pirates/api/request"
	"Pirates/api/response"
	"Pirates/events"
	"Pirates/models/info"
	"Pirates/models/player"
	"Pirates/models/ship"
	"errors"
	"fmt"
)

func (g *Game) SetActionForShipId(shipId string, action request.Action) {
	g.ocean.SetActionForShip(shipId, action)
}

func (g *Game) PlayerHasShip(player *player.Player, shipId string) bool {
	for _, pship := range player.Ships {
		if pship.Deleted {
			continue
		}
		if pship.Id == shipId {
			return true
		}
	}

	return false
}

func (g *Game) GetInfoForShipId(shipId string) *response.Info {
	playerShip := g.ocean.GetShipForId(shipId)
	if playerShip == nil || playerShip.Deleted {
		return nil
	}

	resp := response.Info{}
	resp.Lookout = g.ocean.GetLookoutForShip(shipId)
	resp.Ship = playerShip.GetInfoShip()

	return &resp
}

func (g *Game) GetShipsForPlayer(player *player.Player) ([]*info.InfoShip, error) {
	ships := make([]*info.InfoShip, 0, 0)
	for _, pship := range player.Ships {
		if pship.Deleted {
			continue
		}
		ships = append(ships, pship.GetInfoShip())
	}

	return ships, nil
}

func (g *Game) NewShip(buyRequest *request.Buy) (*info.InfoShip, error) {
	player := g.GetPlayerWithRequest(buyRequest.Player)
	if player == nil {
		return nil, errors.New("[PLAYER-1] Player with Secret not found")
	}

	price := g.calcShipPrice(buyRequest)
	if price > player.Gold {
		return nil, errors.New("[GOLD-1] Not enough Gold")
	}

	player.Gold -= price
	player.GoldSpent += price

	newShip := ship.NewShip(buyRequest.ShipName, buyRequest.Speed, buyRequest.Canons, buyRequest.Sight)
	g.ocean.AddShip(newShip)
	player.Ships = append(player.Ships, newShip)

	events.GetInstance().CreateEvent(
		fmt.Sprintf("[NEW-SHIP] %s bought a new ship with the name %s", buyRequest.Player.Name, buyRequest.ShipName),
		newShip.X,
		newShip.Y,
	)

	return newShip.GetInfoShip(), nil
}

func (g *Game) calcShipPrice(buyRequest *request.Buy) int64 {
	price := 2000
	price += buyRequest.Speed * 1000
	price += buyRequest.Canons * 1000
	price += buyRequest.Sight * 1000

	return int64(price)
}
