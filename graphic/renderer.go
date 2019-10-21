package graphic

import "Pirates/models/ocean"

type Renderer interface {
	Draw(ocean *ocean.Ocean) error
}
