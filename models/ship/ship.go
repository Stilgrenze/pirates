package ship

import (
	"Pirates/api/request"
	"Pirates/events"
	"Pirates/models/info"
	"Pirates/util"
	"fmt"
	"math/rand"
	"time"
)

type Ship struct {
	Id   string
	Name string

	Cannons int
	// Tons deadweight all told
	Tdwat    int
	Hold     int
	Sight    int
	MaxSpeed int

	X int
	Y int

	action info.Action

	Lookout  *info.LookOut
	Messages []info.Message

	NoAction int
	Deleted  bool

	attackShips []*Ship
	Gold        int
}

func NewShip(name string, speed, cannons, sight int) *Ship {
	s := new(Ship)
	s.Id = util.RandSeq(16)
	s.Name = name
	s.Tdwat = 1000

	s.MaxSpeed = 1 + speed
	s.Cannons = 1 + cannons
	s.Sight = 1 + sight
	s.attackShips = make([]*Ship, 0, 0)
	return s
}

func (s *Ship) Tick(lookout *info.LookOut) {
	s.Lookout = lookout
	s.NoAction += 1

	if len(s.attackShips) > 0 {
		for _, attacker := range s.attackShips {
			s.Fight(attacker)
		}
	}
	s.attackShips = make([]*Ship, 0, 0)
}

func (s *Ship) Fight(a *Ship) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	saw := r.Intn(s.Cannons + 1)
	aaw := r.Intn(a.Cannons + 1)

	// Ship Killed
	if aaw >= saw {
		s.Deleted = true

		a.Gold += 1000
		events.GetInstance().CreateEvent(
			fmt.Sprintf("[SHIP-DESTROYED] %s destroyed %s. %s gets 1000 gold as reward", a.Name, s.Name, a.Name),
			a.X,
			a.Y,
		)
	}

	// Only Defender can sink
	if saw > aaw {
		return
	}
}

func (s *Ship) AddAttackShip(ship2 *Ship) {
	// Only attack other ships
	if s.Id != ship2.Id {
		s.attackShips = append(s.attackShips, ship2)
	}
}

func (s *Ship) SetAction(action request.Action) {
	s.NoAction = 0
	s.action = info.Action{
		action.MoveX,
		action.MoveY,
		action.Attack,
	}
}

func (s *Ship) Action() info.Action {
	action := s.action
	s.action = info.Action{}
	return action
}

func (s *Ship) GetDraught() int {
	return int((100.0 / float64(s.Tdwat)) * float64(s.Hold))
}

func (s *Ship) GetSpeed() int {
	return s.MaxSpeed
	/*	speed := int(float64(s.MaxSpeed) * float64(s.GetDraught()/100))
		if speed < 1 {
			return 1
		}
		return speed*/
}

func (s *Ship) AppendMessage(msg info.Message) {
	s.Messages = append(s.Messages, msg)
}

func (s *Ship) GetInfoShip() *info.InfoShip {
	return &info.InfoShip{
		s.Id,
		s.Name,
		s.Sight,
		s.Cannons,
		s.Tdwat,
		s.GetDraught(),
		s.GetSpeed(),
	}
}
