package linkedlist

import (
	"github.com/mstansbu/go-datastructures/nodes"
)

type CircularLinkedList[T comparable] struct {
	Head  *nodes.DLNode[T]
	count int
}

func NewCircularLinkedList[T comparable](values ...T) CircularLinkedList[T] {
	newCircularLinkedList := CircularLinkedList[T]{nil, 0}
	if len(values) != 0 {
		newCircularLinkedList.Append(values...)
	}
	return newCircularLinkedList
}

func (this *CircularLinkedList[T]) Get(position int) (T, error) {
	err := this.positionCheck(position)
	if err != nil {
		var val T
		return val, err
	}
	if this.count/2 > position {
		index := 0
		curr := this.Head
		for index <= position {
			if index == position {
				return curr.Val, nil
			}
			index++
			curr = curr.Next
		}
	} else {
		index := this.count - 1
		curr := this.Head.Prev
		for index >= position {
			if index == position {
				return curr.Val, nil
			}
			index--
			curr = curr.Prev
		}
	}
	panic("idk man, Get went OOB")
}

func (this *CircularLinkedList[T]) Set(position int, value T) error {
	err := this.positionCheck(position)
	if err != nil {
		return err
	}
	if this.count/2 > position {
		index := 0
		curr := this.Head
		for index <= position {
			if index == position {
				curr.Val = value
				return nil
			}
			index++
			curr = curr.Next
		}
	} else {
		index := this.count - 1
		curr := this.Head.Prev
		for index >= position {
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

func (this *CircularLinkedList[T]) GetFirst(value T) (int, error) {
	if this.Empty() {
		return -1, ErrEmptyList
	}
	index := 0
	curr := this.Head
	for index < this.count {
		if curr.Val == value {
			return index, nil
		}
		index++
		curr = curr.Next
	}
	return -1, ErrValueNotFound
}

func (this *CircularLinkedList[T]) GetLast(value T) (int, error) {
	if this.Empty() {
		return -1, ErrEmptyList
	}
	index := this.count - 1
	curr := this.Head.Prev
	for curr != nil {
		if curr.Val == value {
			return index, nil
		}
		index--
		curr = curr.Prev
	}
	return -1, ErrValueNotFound
}

func (this *CircularLinkedList[T]) GetAll(value T) ([]int, error) {
	if this.Empty() {
		return nil, ErrEmptyList
	}
	index := 0
	positions := make([]int, 0)
	curr := this.Head
	for index < this.count {
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

/*func (this *CircularLinkedList[T]) GetAllFirst(values ...T) ([]int, error) {
	indexMap := make(map[T]int)
	for i, value := range values {

	}
}

func (this *CircularLinkedList[T]) GetAllLast(values ...T) ([]int, error) {

}

func (this *CircularLinkedList[T]) GetAllAll(values ...T) ([][]int, error) {

}*/

func (this *CircularLinkedList[T]) Append(values ...T) {
	for _, value := range values {
		node := nodes.NewDLNode(value)
		if this.Empty() {
			this.Head = &node
		} else {
			this.Head.Prev.Next = &node
			node.Prev = this.Head.Prev
			node.Next = this.Head
			this.Head.Prev = &node
		}
		this.count++
	}
}

func (this *CircularLinkedList[T]) Prepend(values ...T) {
	curr := this.Head
	for _, value := range values {
		node := nodes.NewDLNode(value)
		if this.Empty() {
			this.Head = &node
			curr = this.Head
		} else {
			curr.Prev.Next = &node
			node.Prev = curr.Prev
			node.Next = curr
			curr.Prev = &node
			curr = &node
		}
		this.count++
	}
}

func (this *CircularLinkedList[T]) InsertAfter(position int, values ...T) error {
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
		for index <= position {
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
		curr := this.Head
		for index >= position {
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

func (this *CircularLinkedList[T]) InsertBefore(position int, values ...T) error {
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
		for index <= position {
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
		curr := this.Head
		for index >= position {
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

func (this *CircularLinkedList[T]) Empty() bool {
	return this.count == 0
}

func (this *CircularLinkedList[T]) Size() int {
	return this.count
}

/*func (this *CircularLinkedList[T]) FirstIndexOfAll(values ...T) ([]int, bool) {

}

func (this *CircularLinkedList[T]) LastIndexOfAll(values ...T) ([]int, bool) {

}

func (this *CircularLinkedList[T]) AllIndexOfAll(values ...T) ([][]int, bool) {

}*/

func (this *CircularLinkedList[T]) Clear() {
	this.Head = nil
	this.count = 0
}
func (this *CircularLinkedList[T]) doDelete(curr *nodes.DLNode[T]) {
	curr.Prev.Next = curr.Next
	curr.Next.Prev = curr.Prev
	this.count--
}

func (this *CircularLinkedList[T]) Delete(position int) error {
	err := this.positionCheck(position)
	if err != nil {
		return err
	}
	if position == 0 {
		this.Head.Prev.Next = this.Head.Next
		this.Head.Next.Prev = this.Head.Prev
		this.Head = this.Head.Next
		this.count--
		return nil
	}
	if this.count/2 > position {
		index := 1
		curr := this.Head.Next
		for index <= position {
			if index == position {
				this.doDelete(curr)
				return nil
			}
			index++
			curr = curr.Next
		}
	} else {
		index := this.count - 1
		curr := this.Head
		for index >= position {
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

func (this *CircularLinkedList[T]) DeleteFirst(value T) error {
	if this.Empty() {
		return ErrEmptyList
	} else if this.Head.Val == value {
		this.Head.Prev.Next = this.Head.Next
		this.Head.Next.Prev = this.Head.Prev
		this.Head = this.Head.Next
		this.count--
		return nil
	}
	curr := this.Head.Next
	index := 1
	for index < this.count {
		if curr.Val == value {
			this.doDelete(curr)
			return nil
		}
		curr = curr.Next
		index++
	}
	return ErrValueNotFound
}

func (this *CircularLinkedList[T]) DeleteLast(value T) error {
	if this.Empty() {
		return ErrEmptyList
	}
	curr := this.Head.Prev
	index := this.count - 1
	for index >= 0 {
		if curr.Val == value {
			this.doDelete(curr)
			return nil
		}
		curr = curr.Prev
		index--
	}
	return ErrValueNotFound
}

func (this *CircularLinkedList[T]) DeleteAll(value T) error {
	if this.Empty() {
		return ErrEmptyList
	}
	found := false
	curr := this.Head
	index := 0
	lastIndex := this.count - 1
	for index <= lastIndex {
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

/*func (this *CircularLinkedList[T]) DeleteAllFirst(values ...T) {

}

func (this *CircularLinkedList[T]) DeleteAllLast(values ...T) {

}

func (this *CircularLinkedList[T]) DeleteAllAll(values ...T) {

}*/

func (this *CircularLinkedList[T]) Copy() LinkedList[T] {
	return NewLinkedList(this.ToArray()...)
}

func (this *CircularLinkedList[T]) ToArray() []T {
	array := make([]T, 0, this.count)
	curr := this.Head
	index := 0
	for index < this.count {
		array = append(array, curr.Val)
		curr = curr.Next
		index++
	}
	return array
}

/*func (this *CircularLinkedList[T]) search(fn action) {
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

func (this *CircularLinkedList[T]) positionCheck(position int) error {
	if this.Empty() {
		return ErrEmptyList
	} else if position < 0 || this.count <= position {
		return ErrIndexOutOfBounds
	}
	return nil
}
