package ocean

import (
	"Pirates/api/request"
	"Pirates/events"
	"Pirates/models/info"
	"Pirates/models/ship"
	"fmt"
	"github.com/jinzhu/copier"
	"math/rand"
)

const MAX_NO_ACTION = 600
const PORT_PERCENT = 10

type Ocean struct {
	Tiles         [][]Tile
	tilesProgress [][]Tile

	Width  int
	Height int

	tick int64

	Ships map[string]*ship.Ship
	Infos []string
}

func GenerateOcean(width int, height int) *Ocean {
	o := new(Ocean)
	o.Width = width
	o.Height = height
	o.Tiles = make([][]Tile, o.Width)
	o.tilesProgress = make([][]Tile, o.Width)
	o.Ships = make(map[string]*ship.Ship)

	for x := 0; x < len(o.Tiles); x++ {
		o.Tiles[x] = make([]Tile, o.Height)
		o.tilesProgress[x] = make([]Tile, o.Height)

		for y := 0; y < len(o.Tiles[x]); y++ {
			tile := NewTile(PORT_PERCENT)
			tilep := Tile{}
			copier.Copy(&tilep, &tile)

			o.Tiles[x][y] = tile
			o.tilesProgress[x][y] = tilep
		}
	}

	return o
}

func (o *Ocean) AddShip(ship *ship.Ship) {
	o.Ships[ship.Id] = ship

	ship.X = rand.Intn(o.Width)
	ship.Y = rand.Intn(o.Height)

	o.Tiles[ship.X][ship.Y].AddShip(ship)
}

func (o *Ocean) RemoveShip(ship *ship.Ship) {
	ship.Deleted = true
	delete(o.Ships, ship.Id)
}

func (o *Ocean) GetShipForId(shipId string) *ship.Ship {
	if _, ok := o.Ships[shipId]; ok {
		return o.Ships[shipId]
	}
	return nil
}

func (o *Ocean) SetActionForShip(shipId string, action request.Action) {
	if _, ok := o.Ships[shipId]; ok {
		o.Ships[shipId].SetAction(action)
	}
}

func (o *Ocean) GetLookoutForShip(shipId string) *info.LookOut {
	if _, ok := o.Ships[shipId]; ok {
		return o.Ships[shipId].Lookout
	} else {
		return nil
	}
}

func (o *Ocean) Tick() {
	o.startProgress()

	// Handle Tiles first
	for x := 0; x < len(o.Tiles); x++ {
		for y := 0; y < len(o.Tiles[x]); y++ {
			o.handleTile(&o.Tiles[x][y])
		}
	}

	o.endProgress()

	// Update all Tiles
	for x := 0; x < len(o.Tiles); x++ {
		for y := 0; y < len(o.Tiles[x]); y++ {
			tile := &o.Tiles[x][y]
			tile.Tick()

			for i := 0; i < len(tile.Ships); i++ {
				ship := tile.Ships[i]
				lookout := o.getLookOut(ship.X, ship.Y, ship.Sight)
				ship.Tick(lookout)

				// Remove inactive Ships
				if ship.NoAction > MAX_NO_ACTION || ship.Deleted {
					tile.RemoveShip(ship)
					o.RemoveShip(ship)
				}
			}
		}
	}

	o.tick++
}

func (o *Ocean) startProgress() {
	for x := 0; x < len(o.Tiles); x++ {
		for y := 0; y < len(o.Tiles[x]); y++ {
			o.tilesProgress[x][y] = o.Tiles[x][y]
		}
	}
}

func (o *Ocean) endProgress() {
	for x := 0; x < len(o.Tiles); x++ {
		for y := 0; y < len(o.Tiles[x]); y++ {
			o.Tiles[x][y] = o.tilesProgress[x][y]
		}
	}
}

func (o *Ocean) handleTile(tile *Tile) {
	for i := 0; i < len(tile.Ships); i++ {
		ship := tile.Ships[i]
		o.handleShip(tile, ship)
	}
	if tile.Port != nil {
		tile.Port.Tick()
	}
}

func (o *Ocean) handleShip(tile *Tile, ship *ship.Ship) {
	action := ship.Action()

	for _, attackId := range action.Attack {
		// Port attack
		if tile.Port != nil && tile.Port.Id == attackId {
			tile.Port.AddAttackShip(ship)
			events.GetInstance().CreateEvent(
				fmt.Sprintf("[PORT-ATTACK] %s attacking %s", ship.Name, tile.Port.Name),
				ship.X,
				ship.Y,
			)
		}

		// Ship Attack
		for _, aship := range tile.Ships {
			if aship.Id == attackId {
				if aship.Deleted {
					continue
				}
				aship.AddAttackShip(ship)
				events.GetInstance().CreateEvent(
					fmt.Sprintf("[SHIP-ATTACK] %s attacking %s", ship.Name, aship.Name),
					ship.X,
					ship.Y,
				)
			}
		}
	}

	o.moveShip(ship, action.MoveX, action.MoveY)
}

func (o *Ocean) getLookOut(x, y, sight int) *info.LookOut {
	tiles := make([]info.InfoTile, 0, 0)
	for lx := x - sight; lx <= x+sight; lx++ {
		for ly := y - sight; ly <= y+sight; ly++ {
			infoTile := o.getTile(lx, ly)
			if infoTile != nil {
				tiles = append(tiles, infoTile.GetInfo(lx, ly))
			}
		}
	}

	return &info.LookOut{
		x, y, tiles,
	}
}

func (o *Ocean) getTile(x, y int) *Tile {
	if x >= o.Width || x < 0 || y >= o.Height || y < 0 {
		return nil
	}

	return &o.Tiles[x][y]
}

func (o *Ocean) moveShip(ship *ship.Ship, dx int, dy int) {
	if dx+dy > ship.MaxSpeed {
		ship.AppendMessage(info.SPEED_NOT_POSSIBLE)
		return
	}

	o.moveShipOneStep(ship, dx, dy)
}

func (o *Ocean) moveShipOneStep(ship *ship.Ship, dx int, dy int) {
	o.tilesProgress[ship.X][ship.Y].RemoveShip(ship)
	newX := ship.X + dx
	newY := ship.Y + dy

	// Border X
	if newX < 0 {
		newX = 0
	}
	if newX >= o.Width {
		newX = o.Width - 1
	}
	// Border Y
	if newY < 0 {
		newY = 0
	}
	if newY >= o.Height {
		newY = o.Height - 1
	}

	ship.X = newX
	ship.Y = newY

	o.tilesProgress[ship.X][ship.Y].AddShip(ship)
}
