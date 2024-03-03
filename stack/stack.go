package stack

import (
	"github.com/mstansbu/go-datastructures/nodes"
)

type Stack[T comparable] struct {
	Head  *nodes.Node[T]
	Count int
}

func (this *Stack[T]) NewStack(node nodes.Node[T]) Stack[T] {
	nodeTo := &node
	return Stack[T]{nodeTo, 1}
}

func (this *Stack[T]) Pop() *nodes.Node[T] {
	if this.Head == nil {
		return nil
	}
	returnValue := this.Head
	this.Head = this.Head.Next
	this.Count--
	return returnValue
}

func (this *Stack[T]) Peek() *nodes.Node[T] {
	return this.Head
}

func (this *Stack[T]) Push(node *nodes.Node[T]) {
	node.Next = this.Head
	this.Head = node
	this.Count++
}

func (this *Stack[T]) Size() int {
	return this.Count
}

func (this *Stack[T]) Search(value T) bool {
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
