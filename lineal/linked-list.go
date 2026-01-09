package lineal

import (
	"errors"
	"fmt"
)

type node[T comparable] struct {
	value T
	next  *node[T]
}

func NewNode[T comparable](value T) *node[T] {
	return &node[T]{
		value: value,
	}
}

func (node node[T]) GetValue() T {
	return node.value
}

func (node *node[T]) SetValue(value T) {
	node.value = value
}

func (node *node[T]) SetNext(next *node[T]) {
	node.next = next
}

func (node *node[T]) GetNext() *node[T] {
	return node.next
}

type LinkedList[T comparable] struct {
	head *node[T]
}

func NewLinkedList[T comparable]() *LinkedList[T] {
	return &LinkedList[T]{}
}

func (link *LinkedList[T]) SetHead(node *node[T]) {
	link.head = node
}

func (link LinkedList[T]) GetHead() *node[T] {
	return link.head
}

func (link *LinkedList[T]) InsertAtBeginning(value T) {
	newNode := NewNode(value)
	newNode.SetNext(link.GetHead())
	link.SetHead(newNode)
}

func (link *LinkedList[T]) InsertAfter(prevNode *node[T], value T) {
	newNode := NewNode(value)
	newNode.SetNext(prevNode.GetNext())
	prevNode.SetNext(newNode)
}

func (link *LinkedList[T]) InsertAtEnd(value T) {
	newNode := NewNode(value)

	if link.head == nil {
		link.SetHead(newNode)
		return
	}

	last := link.GetHead()

	for last.GetNext() != nil {
		last = last.GetNext()
	}

	last.SetNext(newNode)
}

func (link *LinkedList[T]) DeleteNode(position uint) error {
	if link.GetHead() == nil {
		return errors.New("LinkedList is empty")
	}

	if position == 0 {
		link.SetHead(link.GetHead().GetNext())
		return nil
	}

	temp := link.GetHead()

	for range position - 1 {
		if temp.GetNext() == nil {
			break
		}
		temp = temp.GetNext()
	}

	if temp.GetNext() == nil {
		return errors.New("Position out of index")
	}

	temp.SetNext(temp.GetNext().GetNext())

	return nil
}

func (link *LinkedList[T]) Search(value T) *node[T] {
	current := link.GetHead()

	for current != nil {
		if current.GetValue() == value {
			break
		}
	}

	return &node[T]{}
}

func (link *LinkedList[T]) Traversal() {
	temp := link.GetHead()

	for temp != nil {
		fmt.Println("Value: ", temp.GetValue(), " ", "Pointer: ", temp.GetNext())
		temp = temp.GetNext()
	}
}
