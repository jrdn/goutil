package types

type Set[T any] interface {
	Sliceable[T]
	Lengthable[T]

	Add(val T) Set[T]
	Remove(val T) Set[T]
	Contains(val T) bool

	Intersect(Set[T]) Set[T]
	Union(Set[T]) Set[T]
	Subtract(Set[T]) Set[T]
}

func NewSet[T comparable]() Set[T] {
	return &setImpl[T]{
		data: make(map[T]struct{}),
	}
}

func NewSetFromSlice[T comparable](data []T) Set[T] {
	d := make(map[T]struct{})
	for _, x := range data {
		d[x] = struct{}{}
	}
	return &setImpl[T]{
		data: d,
	}
}

type setImpl[T comparable] struct {
	data map[T]struct{}
}

func (s *setImpl[T]) ToSlice() []T {
	x := make([]T, len(s.data))

	i := 0
	for key := range s.data {
		x[i] = key
		i++
	}

	return x
}

func (s *setImpl[T]) Add(val T) Set[T] {
	s.data[val] = struct{}{}
	return s
}

func (s *setImpl[T]) Remove(val T) Set[T] {
	delete(s.data, val)
	return s
}

func (s *setImpl[T]) Contains(val T) bool {
	_, ok := s.data[val]
	return ok
}

func (s *setImpl[T]) Len() int {
	return len(s.data)
}

func (s *setImpl[T]) Intersect(other Set[T]) Set[T] {
	x := NewSet[T]()
	for _, a := range s.ToSlice() {
		if other.Contains(a) {
			x.Add(a)
		}
	}

	return x
}

func (s *setImpl[T]) Union(other Set[T]) Set[T] {
	x := NewSet[T]()
	for _, a := range s.ToSlice() {
		x.Add(a)
	}

	for _, b := range other.ToSlice() {
		x.Add(b)
	}
	return x
}

func (s *setImpl[T]) Subtract(other Set[T]) Set[T] {
	x := NewSet[T]()
	for _, a := range s.ToSlice() {
		x.Add(a)
	}

	for _, b := range other.ToSlice() {
		x.Remove(b)
	}
	return x
}
