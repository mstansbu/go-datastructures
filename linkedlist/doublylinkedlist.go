package linkedlist

import (
	"github.com/mstansbu/go-datastructures/nodes"
)

type DoublyLinkedList[T comparable] struct {
	Head  *nodes.DLNode[T]
	Tail  *nodes.DLNode[T]
	count int
}

func NewDoublyLinkedList[T comparable](values ...T) DoublyLinkedList[T] {
	newDoublyLinkedList := DoublyLinkedList[T]{nil, nil, 0}
	if len(values) != 0 {
		newDoublyLinkedList.Append(values...)
	}
	return newDoublyLinkedList
}

func (this *DoublyLinkedList[T]) Get(position int) (T, error) {
	err := this.positionCheck(position)
	if err != nil {
		var val T
		return val, err
	}
	if this.count/2 > position {
		index := 0
		curr := this.Head
		for curr != nil {
			if index == position {
				return curr.Val, nil
			}
			index++
			curr = curr.Next
		}
	} else {
		index := this.count - 1
		curr := this.Tail
		for curr != nil {
			if index == position {
				return curr.Val, nil
			}
			index--
			curr = curr.Prev
		}
	}
	panic("idk man, Get went OOB")
}

func (this *DoublyLinkedList[T]) Set(position int, value T) error {
	err := this.positionCheck(position)
	if err != nil {
		return err
	}
	if this.count/2 > position {
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
	} else {
		index := this.count - 1
		curr := this.Tail
		for curr != nil {
			if index == position {
				curr.Val = value
				return nil
			}
			index--
			curr = curr.Prev
		}
	}
	panic("idk man, Get went OOB")
}

func (this *DoublyLinkedList[T]) GetFirst(value T) (int, error) {
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

func (this *DoublyLinkedList[T]) GetLast(value T) (int, error) {
	if this.Empty() {
		return -1, ErrEmptyList
	}
	index := this.count - 1
	curr := this.Tail
	for curr != nil {
		if curr.Val == value {
			return index, nil
		}
		index--
		curr = curr.Prev
	}
	return -1, ErrValueNotFound
}

func (this *DoublyLinkedList[T]) GetAll(value T) ([]int, error) {
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

/*func (this *DoublyLinkedList[T]) GetAllFirst(values ...T) ([]int, error) {
	indexMap := make(map[T]int)
	for i, value := range values {

	}
}

func (this *DoublyLinkedList[T]) GetAllLast(values ...T) ([]int, error) {

}

func (this *DoublyLinkedList[T]) GetAllAll(values ...T) ([][]int, error) {

}*/

func (this *DoublyLinkedList[T]) Append(values ...T) {
	for _, value := range values {
		node := nodes.NewDLNode(value)
		if this.Empty() {
			this.Head, this.Tail = &node, &node
		} else {
			this.Tail.Next = &node
			node.Prev = this.Tail
			this.Tail = this.Tail.Next
		}
		this.count++
	}
}

func (this *DoublyLinkedList[T]) Prepend(values ...T) {
	for _, value := range values {
		node := nodes.NewDLNode(value)
		if this.Empty() {
			this.Head, this.Tail = &node, &node
		} else {
			node.Next = this.Head
			this.Head.Prev = &node
			this.Head = this.Head.Prev
		}
		this.count++
	}
}

func (this *DoublyLinkedList[T]) InsertAfter(position int, values ...T) error {
	err := this.positionCheck(position)
	if err != nil {
		return err
	}
	if position == this.count-1 {
		this.Append(values...)
		return nil
	}
	if this.count/2 > position {
		index := 0
		curr := this.Head
		for curr != nil {
			if index == position {
				for _, value := range values {
					node := nodes.NewDLNode(value)
					curr.Next.Prev = &node
					node.Next = curr.Next
					node.Prev = curr
					curr.Next = &node
					curr = &node
					this.count++
				}
				return nil
			}
			index++
			curr = curr.Next
		}
	} else {
		index := this.count - 1
		curr := this.Tail
		for curr != nil {
			if index == position {
				for _, value := range values {
					node := nodes.NewDLNode(value)
					curr.Next.Prev = &node
					node.Next = curr.Next
					node.Prev = curr
					curr.Next = &node
					curr = &node
					this.count++
				}
				return nil
			}
			index--
			curr = curr.Prev
		}
	}
	panic("idk man, Insert went OOB")
}

func (this *DoublyLinkedList[T]) InsertBefore(position int, values ...T) error {
	err := this.positionCheck(position)
	if err != nil {
		return err
	}
	if position == 0 {
		this.Prepend(values...)
		return nil
	}
	if this.count/2 > position {
		index := 1
		curr := this.Head.Next
		for curr != nil {
			if index == position {
				for _, value := range values {
					node := nodes.NewDLNode(value)
					curr.Prev.Next = &node
					node.Prev = curr.Prev
					node.Next = curr
					curr.Prev = &node
					curr = &node
					this.count++
				}
				return nil
			}
			index++
			curr = curr.Next
		}
	} else {
		index := this.count - 1
		curr := this.Tail
		for curr != nil {
			if index == position {
				for _, value := range values {
					node := nodes.NewDLNode(value)
					curr.Prev.Next = &node
					node.Prev = curr.Prev
					node.Next = curr
					curr.Prev = &node
					curr = &node
					this.count++
				}
				return nil
			}
			index--
			curr = curr.Prev
		}
	}
	panic("idk man, Insert went OOB")
}

func (this *DoublyLinkedList[T]) Empty() bool {
	return this.count == 0
}

func (this *DoublyLinkedList[T]) Size() int {
	return this.count
}

/*func (this *DoublyLinkedList[T]) FirstIndexOfAll(values ...T) ([]int, bool) {

}

func (this *DoublyLinkedList[T]) LastIndexOfAll(values ...T) ([]int, bool) {

}

func (this *DoublyLinkedList[T]) AllIndexOfAll(values ...T) ([][]int, bool) {

}*/

func (this *DoublyLinkedList[T]) Clear() {
	this.Head, this.Tail = nil, nil
	this.count = 0
}
func (this *DoublyLinkedList[T]) doDelete(curr *nodes.DLNode[T]) {
	curr.Prev.Next = curr.Next
	curr.Next.Prev = curr.Prev
	this.count--
}

func (this *DoublyLinkedList[T]) Delete(position int) error {
	err := this.positionCheck(position)
	if err != nil {
		return err
	}
	if position == 0 {
		this.Head = this.Head.Next
		this.count--
		return nil
	}
	if this.count/2 > position {
		index := 1
		curr := this.Head.Next
		for curr != nil {
			if index == position {
				this.doDelete(curr)
				return nil
			}
			index++
			curr = curr.Next
		}
	} else {
		index := this.count - 1
		curr := this.Tail
		for curr != nil {
			if index == position {
				this.doDelete(curr)
				return nil
			}
			index--
			curr = curr.Prev
		}
	}
	panic("idk man, Insert went OOB")
}

func (this *DoublyLinkedList[T]) DeleteFirst(value T) error {
	if this.Empty() {
		return ErrEmptyList
	} else if this.Head.Val == value {
		this.Head = this.Head.Next
		this.count--
		return nil
	}
	curr := this.Head.Next
	for curr != nil {
		if curr.Val == value {
			this.doDelete(curr)
			return nil
		}
		curr = curr.Next
	}
	return ErrValueNotFound
}

func (this *DoublyLinkedList[T]) DeleteLast(value T) error {
	if this.Empty() {
		return ErrEmptyList
	} else if this.Tail.Val == value {
		this.Tail = this.Tail.Prev
		this.count--
		return nil
	}
	curr := this.Head.Next
	for curr != nil {
		if curr.Val == value {
			this.doDelete(curr)
			return nil
		}
		curr = curr.Prev
	}
	return ErrValueNotFound
}

func (this *DoublyLinkedList[T]) DeleteAll(value T) error {
	if this.Empty() {
		return ErrEmptyList
	}
	found := false
	curr := this.Head
	for curr != nil {
		if curr.Val == value {
			found = true
			this.doDelete(curr)
		}
		curr = curr.Next
	}
	if found {
		return nil
	}
	return ErrValueNotFound
}

/*func (this *DoublyLinkedList[T]) DeleteAllFirst(values ...T) {

}

func (this *DoublyLinkedList[T]) DeleteAllLast(values ...T) {

}

func (this *DoublyLinkedList[T]) DeleteAllAll(values ...T) {

}*/

func (this *DoublyLinkedList[T]) Copy() LinkedList[T] {
	return NewLinkedList(this.ToArray()...)
}

func (this *DoublyLinkedList[T]) ToArray() []T {
	array := make([]T, 0, this.count)
	curr := this.Head
	for curr != nil {
		array = append(array, curr.Val)
		curr = curr.Next
	}
	return array
}

/*func (this *DoublyLinkedList[T]) search(fn action) {
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

func (this *DoublyLinkedList[T]) positionCheck(position int) error {
	if this.Empty() {
		return ErrEmptyList
	} else if position < 0 || this.count <= position {
		return ErrIndexOutOfBounds
	}
	return nil
}
