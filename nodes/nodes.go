package nodes

type Node[T comparable] struct {
	Val  T
	Next *Node[T]
}

func (this *Node[T]) NewNode(value T) Node[T] {
	return Node[T]{Val: value, Next: nil}
}
