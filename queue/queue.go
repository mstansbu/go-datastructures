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
	count uint
}

func NewQueueArray[T comparable](values ...T) QueueArray[T] {
	if len(values) == 0 {
		return QueueArray[T]{items: make([]T, 0), head: 0, tail: 0}
	}
	items := make([]T, len(values), len(values)+100)
	items = append(items, values...)
	return QueueArray[T]{items: items, head: 0, tail: uint(len(items))}
}

func (this *QueueArray[T]) Pop() (T, bool) {
	returnValue, ok := this.Peek()
	if ok {
		this.incrementPointer(this.head)
		this.count--
	}
	return returnValue, ok
}

func (this *QueueArray[T]) Peek() (T, bool) {
	if this.count == 0 {
		var val T
		return val, false
	}
	return this.items[this.head], true

}

func (this *QueueArray[T]) Push(values ...T) {
	for _, value := range values {
		this.items[this.tail] = value
		this.incrementPointer(this.tail)
		if this.tail == this.head {
			this.allocateNewArray(len(values) + 100)
		}
	}
}

func (this *QueueArray[T]) Size() uint {
	return this.count
}

func (this *QueueArray[T]) Empty() bool {
	return this.count == 0
}

func (this *QueueArray[T]) Search(value T) bool {
	if this.items[this.head] == value {
		return true
	}
	currPoint := this.head + 1
	for currPoint != this.tail {
		if this.items[currPoint] == value {
			return true
		}
		this.incrementPointer(currPoint)
	}
	return false
}

func (this *QueueArray[T]) allocateNewArray(addLength int) {
	newItems := make([]T, len(this.items), len(this.items)+addLength)
	newItems = append(newItems, this.items[this.head])
	this.incrementPointer(this.head)
	for this.head != this.tail {
		newItems = append(newItems, this.items[this.head])
		this.incrementPointer(this.head)
	}
	this.items = newItems
	this.head = 0
	this.tail = uint(len(this.items))
}

func (this *QueueArray[T]) incrementPointer(pointer uint) {
	if pointer == uint(len(this.items)-1) {
		pointer = 0
	} else {
		pointer++
	}
}
