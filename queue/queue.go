package queue

import (
	"github.com/mstansbu/go-datastructures/nodes"
)

type QueueList[T comparable] struct {
	Head  *nodes.Node[T]
	Tail  *nodes.Node[T]
	count int
}

func NewQueueList[T comparable](values ...T) QueueList[T] {
	newQueue := QueueList[T]{nil, nil, 0}
	if len(values) != 0 {
		newQueue.Push(values...)
	}
	return newQueue
}

func (this *QueueList[T]) Pop() (T, bool) {
	if this.Head == nil {
		this.count = 0
		var val T
		return val, false
	}
	returnValue := this.Head
	this.Head = this.Head.Next
	this.count--
	return returnValue.Val, true
}

func (this *QueueList[T]) Peek() T {
	return this.Head.Val
}

func (this *QueueList[T]) Push(values ...T) {
	for _, value := range values {
		node := nodes.NewNode(value)
		if this.Head == nil {
			this.Head, this.Tail = &node, &node
		} else {
			this.Tail.Next = &node
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

type QueueArray[T comparable] struct {
	items []T
	head  uint
	tail  uint
}
