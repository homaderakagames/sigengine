package vec2

import "cmp"

type Vector[T cmp.Ordered] struct {
	X, Y T
}

func NewVector[T cmp.Ordered](x, y T) *Vector[T] {
	return &Vector[T]{x, y}
}

func (v *Vector[T]) WithX(value T) *Vector[T] {
	v.X = value
	return v
}

func (v *Vector[T]) WithY(value T) *Vector[T] {
	v.Y = value
	return v
}

func (v *Vector[T]) WithXY(x, y T) *Vector[T] {
	v.X = x
	v.Y = y
	return v
}
