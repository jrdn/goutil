package types

type Tree[T any] interface {
	Get() T
	Add(T) Tree[T]
	Children() []Tree[T]
	Set(T) Tree[T]
}

func NewTree[T any]() Tree[T] {
	return &treeNode[T]{}
}

type treeNode[T any] struct {
	val      T
	children []Tree[T]
}

func (t *treeNode[T]) Set(val T) Tree[T] {
	t.val = val
	return t
}

func (t *treeNode[T]) Get() T {
	return t.val
}

func (t *treeNode[T]) Add(val T) Tree[T] {
	node := &treeNode[T]{
		val: val,
	}
	t.children = append(t.children, node)
	return node
}

func (t *treeNode[T]) Children() []Tree[T] {
	return t.children
}
