package linkedlist

import (
	"errors"

	"github.com/mstansbu/go-datastructures/nodes"
)

var ErrEmptyList = errors.New("cannot perform action on empty list")
var ErrIndexOutOfBounds = errors.New("index provided is out of bounds of the list")
var ErrValueNotFound = errors.New("value not found")

type LinkedList[T comparable] struct {
	Head  *nodes.Node[T]
	Tail  *nodes.Node[T]
	count int
}

func NewLinkedList[T comparable](values ...T) LinkedList[T] {
	newLinkedList := LinkedList[T]{nil, nil, 0}
	if len(values) != 0 {
		newLinkedList.Append(values...)
	}
	return newLinkedList
}

func (this *LinkedList[T]) Get(position int) (T, error) {
	if this.Empty() {
		var val T
		return val, ErrEmptyList
	} else if !this.positionCheck(position) {
		var val T
		return val, ErrIndexOutOfBounds
	}
	index := 0
	curr := this.Head
	for curr != nil {
		if index == position {
			return curr.Val, nil
		}
		index++
		curr = curr.Next
	}
	panic("idk man, Get went OOB")
}

func (this *LinkedList[T]) Set(position int, value T) error {
	if this.Empty() {
		return ErrEmptyList
	} else if !this.positionCheck(position) {
		return ErrIndexOutOfBounds
	}
	index := 0
	curr := this.Head
	for curr != nil {
		if index == position {
			curr.Val = value
			return nil
		}
		index++
		curr = curr.Next
	}
	panic("idk man, Set went OOB")
}

func (this *LinkedList[T]) GetFirst(value T) (int, error) {
	if this.Empty() {
		return -1, ErrEmptyList
	}
	index := 0
	curr := this.Head
	for curr != nil {
		if curr.Val == value {
			return index, nil
		}
		index++
		curr = curr.Next
	}
	return -1, ErrValueNotFound
}

func (this *LinkedList[T]) GetLast(value T) (int, error) {
	if this.Empty() {
		return -1, ErrEmptyList
	}
	index := 0
	position := -1
	curr := this.Head
	for curr != nil {
		if curr.Val == value {
			position = index
		}
		index++
		curr = curr.Next
	}
	if position == -1 {
		return position, ErrValueNotFound
	}
	return position, nil
}

func (this *LinkedList[T]) GetAll(value T) ([]int, error) {
	if this.Empty() {
		return nil, ErrEmptyList
	}
	index := 0
	positions := make([]int, 0)
	curr := this.Head
	for curr != nil {
		if curr.Val == value {
			positions = append(positions, index)
		}
		index++
		curr = curr.Next
	}
	if len(positions) == 0 {
		return nil, ErrValueNotFound
	}
	return positions, nil
}

/*func (this *LinkedList[T]) GetAllFirst(values ...T) ([]int, error) {
	indexMap := make(map[T]int)
	for i, value := range values {

	}
}

func (this *LinkedList[T]) GetAllLast(values ...T) ([]int, error) {

}

func (this *LinkedList[T]) GetAllAll(values ...T) ([][]int, error) {

}*/

func (this *LinkedList[T]) Append(values ...T) {
	for _, value := range values {
		node := nodes.NewNode(value)
		if this.Empty() {
			this.Head, this.Tail = &node, &node
		} else {
			this.Tail.Next = &node
			this.Tail = this.Tail.Next
		}
		this.count++
	}
}

func (this *LinkedList[T]) Prepend(values ...T) {
	for _, value := range values {
		node := nodes.NewNode(value)
		if this.Empty() {
			this.Head, this.Tail = &node, &node
		} else {
			node.Next = this.Head
			this.Head = &node
		}
		this.count++
	}
}

func (this *LinkedList[T]) InsertAt(position int, values ...T) error {
	if this.Empty() {
		return ErrEmptyList
	} else if this.positionCheck(position) {
		return ErrIndexOutOfBounds
	}
	index := 0
	curr := this.Head
	var prev *nodes.Node[T]
	for curr != nil {
		if index == position {
			for _, value := range values {
				node := nodes.NewNode(value)
				node.Next = curr
				prev.Next = &node
				prev = &node
				this.count++
			}
			return nil
		}
		index++
		prev = curr
		curr = curr.Next
	}
	panic("idk man, Insert went OOB")
}

func (this *LinkedList[T]) Empty() bool {
	return this.count == 0
}

func (this *LinkedList[T]) Size() int {
	return this.count
}

/*func (this *LinkedList[T]) FirstIndexOfAll(values ...T) ([]int, bool) {

}

func (this *LinkedList[T]) LastIndexOfAll(values ...T) ([]int, bool) {

}

func (this *LinkedList[T]) AllIndexOfAll(values ...T) ([][]int, bool) {

}*/

func (this *LinkedList[T]) Clear() {
	this.Head, this.Tail = nil, nil
	this.count = 0
}

func (this *LinkedList[T]) Delete(position int) error {
	if this.Empty() {
		return ErrEmptyList
	} else if !this.positionCheck(position) {
		return ErrIndexOutOfBounds
	} else if position == 0 {
		this.Head = this.Head.Next
		return nil
	}
	index := 0
	curr := this.Head
	prev := curr.Next
	for curr != nil {
		if index == position {
			prev.Next = curr.Next
			this.count--
			return nil
		}
		index++
		prev = curr
		curr = curr.Next
	}
	panic("idk man, Insert went OOB")
}

func (this *LinkedList[T]) DeleteFirst(value T) error {
	if this.Empty() {
		return ErrEmptyList
	} else if this.Head.Val == value {
		this.Head = this.Head.Next
		return nil
	}
	curr := this.Head.Next
	prev := this.Head
	for curr != nil {
		if curr.Val == value {
			prev.Next = curr.Next
			return nil
		}
		prev = curr
		curr = curr.Next
	}
	return ErrValueNotFound
}

func (this *LinkedList[T]) DeleteLast(value T) error {
	if this.Empty() {
		return ErrEmptyList
	}
	curr := this.Head.Next
	prev := this.Head
	var currLast *nodes.Node[T]
	var currLastPrev *nodes.Node[T]
	if this.Head.Val == value {
		currLast = this.Head
	}
	for curr != nil {
		if curr.Val == value {
			currLast = curr
			currLastPrev = prev
		}
		prev = curr
		curr = curr.Next
	}
	if currLast != nil {
		if currLastPrev != nil {
			currLastPrev.Next = currLast.Next
		} else {
			this.Head = this.Head.Next
		}
		return nil
	}
	return ErrValueNotFound
}

func (this *LinkedList[T]) DeleteAll(value T) error {
	if this.Empty() {
		return ErrEmptyList
	}
	found := false
	curr := this.Head
	var prev *nodes.Node[T]
	for curr != nil {
		if curr.Val == value {
			found = true
			if prev != nil {
				prev.Next = curr.Next
			} else {
				this.Head = this.Head.Next
			}
		} else {
			prev = curr
		}
		curr = curr.Next
	}
	if found {
		return nil
	}
	return ErrValueNotFound
}

/*func (this *LinkedList[T]) DeleteAllFirst(values ...T) {

}

func (this *LinkedList[T]) DeleteAllLast(values ...T) {

}

func (this *LinkedList[T]) DeleteAllAll(values ...T) {

}*/

func (this *LinkedList[T]) Copy() LinkedList[T] {
	return NewLinkedList(this.ToArray()...)
}

func (this *LinkedList[T]) ToArray() []T {
	array := make([]T, 0, this.count)
	curr := this.Head
	for curr != nil {
		array = append(array, curr.Val)
		curr = curr.Next
	}
	return array
}

/*func (this *LinkedList[T]) search(fn action) {
	if this.Empty() {
		var val T
		return val, ErrEmptyList
	} else if position >= this.count {
		var val T
		return val, ErrIndexOutOfBounds
	}
	index := 0
	curr := this.Head
	for curr != nil {
		if action() {

		}
		if index == position {
			return curr.Val, nil
		}
		index++
	}
}*/

//func Iter() iter.Seq[Item]

func (this *LinkedList[T]) positionCheck(position int) bool {
	return position >= 0 && this.count > position
}
