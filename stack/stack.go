package stack

import (
	"github.com/mstansbu/go-datastructures/nodes"
)

type StackList[T comparable] struct {
	Head  *nodes.Node[T]
	count uint
}

func NewStackList[T comparable](node nodes.Node[T]) StackList[T] {
	nodeTo := &node
	return StackList[T]{nodeTo, 1}
}

func (this *StackList[T]) Pop() *nodes.Node[T] {
	if this.Head == nil {
		this.count = 0
		return nil
	}
	returnValue := this.Head
	this.Head = this.Head.Next
	this.count--
	return returnValue
}

func (this *StackList[T]) Peek() *nodes.Node[T] {
	return this.Head
}

func (this *StackList[T]) Push(nodes ...*nodes.Node[T]) {
	for _, node := range nodes {
		node.Next = this.Head
		this.Head = node
		this.count++
	}
}

func (this *StackList[T]) Size() uint {
	return this.count
}

func (this *StackList[T]) Empty() bool {
	return this.Head == nil
}

func (this *StackList[T]) Search(value T) bool {
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

type StackArray[T comparable] struct {
	items []T
	top   uint
}

func NewStackArray[T comparable](initialSize uint) StackArray[T] {
	stackSize := uint(1000)
	if initialSize != 0 {
		stackSize = initialSize
	}
	return StackArray[T]{items: make([]T, uint(stackSize)), top: 0}
}

func NewStackArrayFromArray[T comparable](values []T) StackArray[T] {
	items := make([]T, len(values), len(values)+100)
	items = append(items, values...)
	return StackArray[T]{items: items, top: uint(len(items))}
}

func (this *StackArray[T]) Pop() (T, bool) {
	returnValue, ok := this.Peek()
	if ok {
		this.top--
		return returnValue, true
	}
	return returnValue, false
}

func (this *StackArray[T]) Peek() (T, bool) {
	if this.top == 0 {
		return this.items[this.top], false
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
