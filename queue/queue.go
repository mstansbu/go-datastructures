package queue

import (
	"errors"

	"github.com/mstansbu/go-datastructures/nodes"
)

var ErrEmptyQueue error = errors.New("cannot perform action on empty queue")

type QueueList[T comparable] struct {
	head  *nodes.Node[T]
	tail  *nodes.Node[T]
	count int
}

func NewQueueList[T comparable](values ...T) QueueList[T] {
	newQueue := QueueList[T]{nil, nil, 0}
	if len(values) != 0 {
		newQueue.Push(values...)
	}
	return newQueue
}

func (this *QueueList[T]) Pop() (T, error) {
	returnValue, ok := this.Peek()
	if ok {
		this.head = this.head.Next
		this.count--
		return returnValue, nil
	}
	return returnValue, ErrEmptyQueue
}

func (this *QueueList[T]) Peek() (T, bool) {
	if this.head == nil {
		this.count = 0
		var val T
		return val, false
	}
	return this.head.Val, true
}

func (this *QueueList[T]) Push(values ...T) {
	for _, value := range values {
		node := nodes.NewNode(value)
		if this.head == nil {
			this.head, this.tail = &node, &node
		} else {
			this.tail.Next = &node
			this.tail = this.tail.Next
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
	curr := this.head
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

func (this *QueueArray[T]) Pop() (T, error) {
	returnValue, ok := this.Peek()
	if ok {
		this.head = this.incrementPointer(this.head)
		this.count--
		return returnValue, nil
	}
	return returnValue, ErrEmptyQueue
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
		this.tail = this.incrementPointer(this.tail)
		if this.tail == this.head {
			this.allocateNewArray(len(values) + 100)
		}
		this.count++
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
		currPoint = this.incrementPointer(currPoint)
	}
	return false
}

func (this *QueueArray[T]) allocateNewArray(addLength int) {
	newItems := make([]T, len(this.items), len(this.items)+addLength)
	newItems = append(newItems, this.items[this.head])
	this.head = this.incrementPointer(this.head)
	for this.head != this.tail {
		newItems = append(newItems, this.items[this.head])
		this.head = this.incrementPointer(this.head)
	}
	this.items = newItems
	this.head = 0
	this.tail = uint(len(this.items))
}

func (this *QueueArray[T]) incrementPointer(pointer uint) uint {
	if pointer == uint(len(this.items)-1) {
		pointer = 0
	} else {
		pointer++
	}
	return pointer
}
