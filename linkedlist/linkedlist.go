package linkedlist

import "github.com/mstansbu/go-datastructures/nodes"

type LinkedList[T comparable] struct {
	Head *nodes.Node[T]
}
