package camera

import (
	"github.com/homaderakagames/sigengine/pkg/engine/gameapp/position"
)

type Camera struct {
	Pos    position.Position[float64]
	Width  float64
	Height float64
}
