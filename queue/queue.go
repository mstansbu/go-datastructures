package queue

import (
	"github.com/mstansbu/go-datastructures/nodes"
)

type Queue[T comparable] struct {
	Head  *nodes.Node[T]
	Tail  *nodes.Node[T]
	Count int
}

func (this *Queue[T]) NewQueue(node nodes.Node[T]) Queue[T] {
	nodeTo := &node
	return Queue[T]{nodeTo, nodeTo, 1}
}

func (this *Queue[T]) Pop() *nodes.Node[T] {
	if this.Head == nil {
		return nil
	}
	returnValue := this.Head
	this.Head = this.Head.Next
	this.Count--
	return returnValue
}

func (this *Queue[T]) Peek() *nodes.Node[T] {
	return this.Head
}

func (this *Queue[T]) Push(node *nodes.Node[T]) {
	if this.Head == nil {
		this.Head, this.Tail = node, node
	} else {
		this.Tail.Next = node
		this.Tail = this.Tail.Next
	}
	this.Count++
}

func (this *Queue[T]) Size() int {
	return this.Count
}

func (this *Queue[T]) Search(value T) bool {
	curr := this.Head
	for curr != nil {
		if curr.Val == value {
			return true
		} else {
			curr = curr.Next
		}
	}
	return false
}
