package vec2

type Rectangle[T Number] struct {
	Min *Vector[T]
	Max *Vector[T]
}

func NewRectangle[T Number](min *Vector[T], max *Vector[T]) *Rectangle[T] {
	rect := &Rectangle[T]{
		Min: min,
		Max: max,
	}
	return rect
}

func (r *Rectangle[T]) Dx() T {
	maxV := NewVector(r.Max.X, r.Max.Y)
	maxV.WithSub(r.Min)
	return maxV.X
}

func (r *Rectangle[T]) Dy() T {
	maxV := NewVector(r.Max.X, r.Max.Y)
	maxV.WithSub(r.Min)
	return maxV.Y
}
