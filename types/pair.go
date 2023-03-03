package types

type Pair[A, B any] interface {
	First() A
	Second() B
	Set(A, B) Pair[A, B]
}

type Triplet[A, B, C any] interface {
	First() A
	Second() B
	Third() C
	Set(A, B, C) Triplet[A, B, C]
}

type Quadruplet[A, B, C, D any] interface {
	First() A
	Second() B
	Third() C
	Fourth() D
	Set(A, B, C, D) Quadruplet[A, B, C, D]
}

func NewPair[A, B any]() Pair[A, B] {
	return &pairImpl[A, B]{}
}

type pairImpl[A, B any] struct {
	first  A
	second B
}

func (p *pairImpl[A, B]) Set(a A, b B) Pair[A, B] {
	p.first = a
	p.second = b
	return p
}

func (p *pairImpl[A, B]) First() A {
	return p.first
}

func (p *pairImpl[A, B]) Second() B {
	return p.second
}

func NewTriplet[A, B, C any]() Triplet[A, B, C] {
	return &tripletImpl[A, B, C]{}
}

type tripletImpl[A, B, C any] struct {
	*pairImpl[A, B]
	third C
}

func (t *tripletImpl[A, B, C]) Set(a A, b B, c C) Triplet[A, B, C] {
	t.first = a
	t.second = b
	t.third = c
	return t
}

func (t *tripletImpl[A, B, C]) Third() C {
	return t.third
}

func NewQuadruplet[A, B, C, D any]() Quadruplet[A, B, C, D] {
	return &quadrupletImpl[A, B, C, D]{}
}

type quadrupletImpl[A, B, C, D any] struct {
	*tripletImpl[A, B, C]
	fourth D
}

func (q *quadrupletImpl[A, B, C, D]) Fourth() D {
	return q.fourth
}

func (q *quadrupletImpl[A, B, C, D]) Set(a A, b B, c C, d D) Quadruplet[A, B, C, D] {
	q.first = a
	q.second = b
	q.third = c
	q.fourth = d
	return q
}
