package stack

import (
	"errors"

	"github.com/mstansbu/go-datastructures/nodes"
)

var ErrEmptyStack error = errors.New("stack is empty, cannot perform this operation")

type StackList[T comparable] struct {
	head  *nodes.Node[T]
	count uint
}

func NewStackList[T comparable](values ...T) StackList[T] {
	newStack := StackList[T]{nil, 0}
	if len(values) != 0 {
		newStack.Push(values...)
	}
	return newStack
}

func (this *StackList[T]) Pop() (T, error) {
	returnValue, ok := this.Peek()
	if ok {
		this.head = this.head.Next
		this.count--
		return returnValue, nil
	}
	return returnValue, ErrEmptyStack
}

func (this *StackList[T]) Peek() (T, bool) {
	if this.head == nil {
		var val T
		return val, false
	}
	return this.head.Val, true
}

func (this *StackList[T]) Push(values ...T) {
	for _, value := range values {
		node := nodes.NewNode(value)
		node.Next = this.head
		this.head = &node
		this.count++
	}
}

func (this *StackList[T]) Size() uint {
	return this.count
}

func (this *StackList[T]) Empty() bool {
	return this.count == 0
}

func (this *StackList[T]) Search(value T) bool {
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

type StackArray[T comparable] struct {
	items []T
	top   uint
}

func NewStackArray[T comparable](values ...T) StackArray[T] {
	if len(values) == 0 {
		return StackArray[T]{items: make([]T, 0), top: 0}
	}
	items := make([]T, len(values), len(values)+100)
	items = append(items, values...)
	return StackArray[T]{items: items, top: uint(len(items))}
}

func (this *StackArray[T]) Pop() (T, error) {
	returnValue, ok := this.Peek()
	if ok {
		this.top--
		return returnValue, nil
	}
	return returnValue, ErrEmptyStack
}

func (this *StackArray[T]) Peek() (T, bool) {
	if this.top == 0 {
		var val T
		return val, false
	}
	return this.items[this.top-1], true
}

func (this *StackArray[T]) Push(values ...T) {
	for _, value := range values {
		this.items[this.top] = value
		this.top++
	}
}

func (this *StackArray[T]) Size() uint {
	return this.top
}

func (this *StackArray[T]) Empty() bool {
	return this.top == 0
}

func (this *StackArray[T]) Search(value T) bool {
	if this.top == 0 {
		return false
	}
	for i := int(this.top) - 1; i >= 0; i-- {
		if this.items[i] == value {
			return true
		}
	}
	return false
}
