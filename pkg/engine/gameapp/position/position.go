package position

import (
	"cmp"
	"github.com/homaderakagames/sigengine/pkg/engine/gameapp/vec2"
)

type Position[T cmp.Ordered] struct {
	vec2.Vector[T]
}
