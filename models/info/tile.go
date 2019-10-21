package info

type TileType string

const (
	WATER TileType = "O"
	SHIP  TileType = "S"
	PORT  TileType = "P"
)

type InfoTile struct {
	X     int
	Y     int
	Type  TileType
	Ships []InfoShip
	Port  *InfoPort
}
