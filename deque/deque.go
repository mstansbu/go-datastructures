package deque

import (
	"errors"

	"github.com/mstansbu/go-datastructures/nodes"
)

var ErrEmptyDeque error = errors.New("cannot perform action on empty deque")

type DequeList[T comparable] struct {
	head  *nodes.DLNode[T]
	tail  *nodes.DLNode[T]
	count int
}

func NewDeque[T comparable](values ...T) DequeList[T] {
	newDeque := DequeList[T]{nil, nil, 0}
	if len(values) != 0 {
		newDeque.PushBack(values...)
	}
	return newDeque
}

func (this *DequeList[T]) PushFront(values ...T) {
	for _, value := range values {
		node := nodes.NewDLNode(value)
		node.Next = this.head
		if this.head != nil {
			this.head.Prev = &node
		}
		this.head = &node
		if this.tail == nil {
			this.tail = &node
		} else if this.tail.Prev == nil {
			this.tail.Prev = &node
		}
		this.count++
	}
}

func (this *DequeList[T]) PushBack(values ...T) {
	for _, value := range values {
		node := nodes.NewDLNode(value)
		node.Next = this.tail
		if this.tail != nil {
			this.tail.Next = &node
		}
		this.tail = &node
		if this.head == nil {
			this.head = &node
		} else if this.head.Next == nil {
			this.head.Next = &node
		}
		this.count++
	}
}

func (this *DequeList[T]) PopFront() (T, error) {
	returnValue, ok := this.PeekFront()
	if ok {
		this.head = this.head.Next
		this.head.Prev = nil
		if this.tail.Prev == nil {
			this.tail = nil
		}
		this.count--
		return returnValue, nil
	}
	return returnValue, ErrEmptyDeque
}

func (this *DequeList[T]) PopBack() (T, error) {
	returnValue, ok := this.PeekBack()
	if ok {
		this.tail = this.tail.Prev
		this.tail.Next = nil
		if this.head.Next == nil {
			this.head = nil
		}
		this.count--
		return returnValue, nil
	}
	return returnValue, ErrEmptyDeque
}

func (this *DequeList[T]) PeekFront() (T, bool) {
	if this.head == nil {
		var val T
		return val, false
	}
	return this.head.Val, true
}

func (this *DequeList[T]) PeekBack() (T, bool) {
	if this.tail == nil {
		var val T
		return val, false
	}
	return this.tail.Val, true
}

func (this *DequeList[T]) Size() int {
	return this.count
}

func (this *DequeList[T]) Empty() bool {
	return this.count == 0
}

func (this *DequeList[T]) Search(value T) bool {
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

type DequeArray[T comparable] struct {
	items []T
	head  uint
	tail  uint
	count uint
}

func NewDequeArray[T comparable](values ...T) DequeArray[T] {
	if len(values) == 0 {
		return DequeArray[T]{items: make([]T, 0), head: 0, tail: 0}
	}
	items := make([]T, len(values), len(values)+100)
	items = append(items, values...)
	return DequeArray[T]{items: items, head: 0, tail: uint(len(items)) - 1}
}

func (this *DequeArray[T]) PopFront() (T, bool) {
	returnValue, ok := this.PeekFront()
	if ok {
		this.head = this.incrementPointer(this.head)
		this.count--
	}
	return returnValue, ok
}

func (this *DequeArray[T]) PeekFront() (T, bool) {
	if this.count == 0 {
		var val T
		return val, false
	}
	return this.items[this.head], true
}

func (this *DequeArray[T]) PopBack() (T, bool) {
	returnValue, ok := this.PeekBack()
	if ok {
		this.tail = this.decrementPointer(this.tail)
		this.count--
	}
	return returnValue, ok
}

func (this *DequeArray[T]) PeekBack() (T, bool) {
	if this.count == 0 {
		var val T
		return val, false
	}
	return this.items[this.tail], true
}

func (this *DequeArray[T]) PushFront(values ...T) {
	for _, value := range values {
		this.head = this.decrementPointer(this.head)
		this.items[this.head] = value
		if this.checkFull() {
			this.allocateNewArray(len(values) + 100)
		}
		this.count++
	}
}

func (this *DequeArray[T]) PushBack(values ...T) {
	for _, value := range values {
		this.tail = this.incrementPointer(this.tail)
		this.items[this.tail] = value
		if this.checkFull() {
			this.allocateNewArray(len(values) + 100)
		}
		this.count++
	}
}

func (this *DequeArray[T]) Size() uint {
	return this.count
}

func (this *DequeArray[T]) Empty() bool {
	return this.count == 0
}

func (this *DequeArray[T]) Search(value T) bool {
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

func (this *DequeArray[T]) allocateNewArray(addLength int) {
	newItems := make([]T, len(this.items), len(this.items)+addLength)
	for this.head != this.tail {
		newItems = append(newItems, this.items[this.head])
		this.head = this.incrementPointer(this.head)
	}
	this.items = newItems
	this.head = 0
	this.tail = uint(len(this.items))
}

func (this *DequeArray[T]) incrementPointer(pointer uint) uint {
	if pointer == uint(len(this.items)-1) {
		pointer = 0
	} else {
		pointer++
	}
	return pointer
}
func (this *DequeArray[T]) decrementPointer(pointer uint) uint {
	if pointer == 0 {
		pointer = uint(len(this.items) - 1)
	} else {
		pointer--
	}
	return pointer
}

func (this *DequeArray[T]) checkFull() bool {
	return (this.head != 0 && this.tail == this.head-1) || (this.head == 0 && this.tail == uint(len(this.items)-1))
}
