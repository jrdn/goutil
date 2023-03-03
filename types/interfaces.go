package types

type Sliceable[T any] interface {
	ToSlice() []T
}

type Lengthable[T any] interface {
	Len() int
}
