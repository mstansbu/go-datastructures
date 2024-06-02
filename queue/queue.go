package queue

import (
	"github.com/mstansbu/go-datastructures/nodes"
)

type QueueList[T comparable] struct {
	Head  *nodes.Node[T]
	Tail  *nodes.Node[T]
	count int
}

func NewQueueList[T comparable](node nodes.Node[T]) QueueList[T] {
	nodeTo := &node
	return QueueList[T]{nodeTo, nodeTo, 1}
}

func (this *QueueList[T]) Pop() *nodes.Node[T] {
	if this.Head == nil {
		return nil
	}
	returnValue := this.Head
	this.Head = this.Head.Next
	this.count--
	return returnValue
}

func (this *QueueList[T]) Peek() *nodes.Node[T] {
	return this.Head
}

func (this *QueueList[T]) Push(nodes ...*nodes.Node[T]) {
	for _, node := range nodes {
		if this.Head == nil {
			this.Head, this.Tail = node, node
		} else {
			this.Tail.Next = node
			this.Tail = this.Tail.Next
		}
		this.count++
	}
}

func (this *QueueList[T]) Size() int {
	return this.count
}

func (this *QueueList[T]) Empty() bool {
	return this.count == 0
}

func (this *QueueList[T]) Search(value T) bool {
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
