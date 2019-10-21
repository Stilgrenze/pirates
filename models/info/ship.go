package info

type InfoShip struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	Sight   int    `json:"sight"`
	Canons  int    `json:"canons"`
	Tdwat   int    `json:"-"`
	Draught int    `json:"-"`
	Speed   int    `json:"speed"`
}
