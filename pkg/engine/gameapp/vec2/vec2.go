package vec2

import "cmp"

type Vector[T cmp.Ordered] struct {
	X, Y T
}
