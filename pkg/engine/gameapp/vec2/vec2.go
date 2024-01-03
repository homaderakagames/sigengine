package vec2

// Number любое число
type Number interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64
}

type Vector[T Number] struct {
	X, Y T
}

func NewVector[T Number](x, y T) *Vector[T] {
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

// WithAdd прибавляет значения вектора o, возвращает указатель на исходный вектор
func (v *Vector[T]) WithAdd(o *Vector[T]) *Vector[T] {
	v.X += o.X
	v.Y += o.Y
	return v
}

func (v *Vector[T]) Equal(o *Vector[T]) bool {
	return v.X == o.X && v.Y == o.Y
}

func (v *Vector[T]) WithSub(other *Vector[T]) *Vector[T] {
	v.X -= other.X
	v.Y -= other.Y
	return v
}
