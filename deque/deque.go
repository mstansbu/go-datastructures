package deque

import "github.com/mstansbu/go-datastructures/nodes"

type Deque[T comparable] struct {
	Head  *nodes.DLNode[T]
	Tail  *nodes.DLNode[T]
	Count int
}

func NewDeque[T comparable](node *nodes.DLNode[T]) Deque[T] {
	return Deque[T]{node, node, 1}
}

func (this *Deque[T]) PushFront(node *nodes.DLNode[T]) {
	node.Next = this.Head
	this.Head.Prev = node
	this.Head = node
}

func (this *Deque[T]) PushBack(node *nodes.DLNode[T]) {
	this.Tail.Next = node
	node.Prev = this.Tail
	this.Tail = node
}

func (this *Deque[T]) PopFront(node *nodes.DLNode[T]) *nodes.DLNode[T] {
	returnValue := this.Head
	this.Head = this.Head.Next
	this.Head.Prev = nil
	return returnValue
}

func (this *Deque[T]) PopBack(node *nodes.DLNode[T]) *nodes.DLNode[T] {
	returnValue := this.Tail
	this.Tail = this.Tail.Prev
	this.Tail.Next = nil
	return returnValue
}

func (this *Deque[T]) PeekFront() *nodes.DLNode[T] {
	return this.Head
}

func (this *Deque[T]) PeekBack() *nodes.DLNode[T] {
	return this.Tail
}
