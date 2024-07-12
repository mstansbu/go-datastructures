package linkedlist

import "github.com/mstansbu/go-datastructures/nodes"

type UnrolledLinkedList[T comparable] struct {
	Head        *nodes.ULNode[T]
	maxElements int
	totalCount  int
}

func NewUnrolledLinkedList[T comparable](maxElements int, values ...T) UnrolledLinkedList[T] {
	newUnrolledLinkedList := UnrolledLinkedList[T]{nil, maxElements, 0}
	if len(values) != 0 {
		newUnrolledLinkedList.Append(values...)
	}
	return newUnrolledLinkedList
}

func (this *UnrolledLinkedList[T]) Append(values ...T) {
	curr := this.Head
	if this.totalCount != 0 {
		for curr.Next != nil {
			curr = curr.Next
		}
	} else {
		node := nodes.NewULNode[T](this.maxElements)
		this.Head, curr = &node, &node
	}

}
