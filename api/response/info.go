package response

import "Pirates/models/info"

type Info struct {
	Lookout *info.LookOut
	Ship    *info.InfoShip
	Error   string `json:"error,omitempty"`
}
