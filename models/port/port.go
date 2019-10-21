package port

import (
	"Pirates/events"
	"Pirates/models/ship"
	"Pirates/util"
	"fmt"
	"github.com/goombaio/namegenerator"
	"math/rand"
	"time"
)

type Port struct {
	Id          string
	Name        string
	Gold        int
	Cannons     int
	GoldPerTick int

	Looted  bool
	Outtime int

	attackShips []*ship.Ship
}

func NewPort(gold, cannons, goldPerTick int) *Port {
	seed := time.Now().UTC().UnixNano()
	nameGenerator := namegenerator.NewNameGenerator(seed)

	name := nameGenerator.Generate()

	port := Port{
		"PORT_" + util.RandSeq(16),
		name,
		gold,
		cannons,
		goldPerTick,
		false,
		0,
		make([]*ship.Ship, 0, 0),
	}
	return &port
}

func (p *Port) AddAttackShip(ship2 *ship.Ship) {
	p.attackShips = append(p.attackShips, ship2)
}

func (p *Port) Tick() {
	if p.Outtime > 0 {
		p.Outtime -= 1
		p.attackShips = make([]*ship.Ship, 0, 0)
		return
	}
	p.Looted = false
	p.Gold += p.GoldPerTick

	gold := p.Gold
	attackShips := len(p.attackShips)
	if attackShips > 0 {
		for _, attacker := range p.attackShips {
			if p.Fight(attacker) {
				treasue := gold / attackShips
				attacker.Gold += treasue
				p.Gold -= treasue
				p.Looted = true

				events.GetInstance().CreateEvent(
					fmt.Sprintf("[PORT-LOOTED] %s looted %d gold in %s", attacker.Name, treasue, p.Name),
					attacker.X,
					attacker.Y,
				)
			}
		}
		p.attackShips = make([]*ship.Ship, 0, 0)
	}
	if p.Looted {
		p.Outtime = rand.Intn(240) + 60
	}
}

func (p *Port) Fight(a *ship.Ship) bool {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	saw := r.Intn(p.Cannons + 1)
	aaw := r.Intn(a.Cannons + 1)

	// Ship Gets Port Something
	if aaw >= saw {
		return true
	}

	// On Equal do Nothing
	if saw > aaw {
		return false
	}

	return false
}
