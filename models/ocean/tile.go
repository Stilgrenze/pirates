package ocean

import (
	"Pirates/models/info"
	"Pirates/models/port"
	"Pirates/models/ship"
	"log"
	"math/rand"
	"time"
)

type Tile struct {
	Type  info.TileType
	Ships []*ship.Ship
	Port  *port.Port
}

func NewTile(portPercentage int) Tile {
	tile := Tile{info.WATER, make([]*ship.Ship, 0), nil}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	if r.Intn(100) < portPercentage {
		tile.addPort()
	}

	return tile
}

func (t *Tile) addPort() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	t.Port = port.NewPort(
		r.Intn(100)+10,
		r.Intn(10),
		r.Intn(5)+1,
	)
}

func (t *Tile) checkType() {
	if t.Port != nil {
		t.Type = info.PORT
	} else if len(t.Ships) > 0 {
		t.Type = info.SHIP
	} else {
		t.Type = info.WATER
	}
}

func (t *Tile) Tick() {
	t.checkType()
}

func (t *Tile) GetInfo(x, y int) info.InfoTile {
	var port *info.InfoPort
	if t.Port != nil {
		port = &info.InfoPort{
			t.Port.Name,
			t.Port.Id,
			t.Port.Gold,
			t.Port.Cannons,
		}
	}

	infoTile := info.InfoTile{x, y, t.Type, t.GetInfoShips(), port}
	return infoTile
}

func (t *Tile) GetInfoShips() []info.InfoShip {
	infoShips := make([]info.InfoShip, 0, 0)
	for i := 0; i < len(t.Ships); i++ {
		ship := t.Ships[i]
		infoShips = append(infoShips, *ship.GetInfoShip())
	}
	return infoShips
}

func (t *Tile) RemoveShip(ship *ship.Ship) {
	for i := 0; i < len(t.Ships); i++ {
		if t.Ships[i].Id == ship.Id {
			t.Ships = append(t.Ships[:i], t.Ships[i+1:]...)
			return
		}
	}
	log.Println("Ship not on Ocean")
}

func (t *Tile) AddShip(ship *ship.Ship) {
	t.Ships = append(t.Ships, ship)
}
