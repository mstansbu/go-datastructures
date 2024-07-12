package nodes

type Node[T comparable] struct {
	Val  T
	Next *Node[T]
}

func NewNode[T comparable](value T) Node[T] {
	return Node[T]{Val: value, Next: nil}
}

type Tuple[T any, N any] struct {
	First  T
	Second N
}

func NewTuple[T, N any](t T, n N) Tuple[T, N] {
	return Tuple[T, N]{t, n}
}

type DLNode[T comparable] struct {
	Val  T
	Next *DLNode[T]
	Prev *DLNode[T]
}

func NewDLNode[T comparable](value T) DLNode[T] {
	return DLNode[T]{Val: value, Next: nil, Prev: nil}
}

type ULNode[T comparable] struct {
	Next      *ULNode[T]
	ItemCount int
	Items     []T
}

func NewULNode[T comparable](maxElements int) ULNode[T] {
	values := make([]T, 0, maxElements)
	return ULNode[T]{Next: nil, ItemCount: 0, Items: values}
}
