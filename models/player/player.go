package player

import "Pirates/models/ship"

type Player struct {
	Name   string
	Secret string

	Gold      int64
	GoldSpent int64

	Ships []*ship.Ship `json:"-"`
}
