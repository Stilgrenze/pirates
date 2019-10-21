package request

type Buy struct {
	Player PlayerRequest

	ShipName string
	Canons   int
	Sight    int
	Speed    int
}
