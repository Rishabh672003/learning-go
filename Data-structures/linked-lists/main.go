package main

import "fmt"

type Node[T any] struct {
	value T
	next  *Node[T]
}

func (n Node[T]) print() {
	fmt.Print(n.value, " ")
}

type Lili[T any] struct {
	head *Node[T]
}

func (l *Lili[T]) print() {
	node := l.head

	for node != nil {
		node.print()
		node = node.next
	}
}

func (l *Lili[T]) addAtEnd(elem T) {
	no := &Node[T]{elem, nil}
	if l.head == nil {
		l.head = no
	} else {
		node := l.head

		for node.next != nil {
			node = node.next
		}
		node.next = no
	}
}

func main() {
	li := Lili[int64]{}
	li.addAtEnd(12)
	li.addAtEnd(12)
	li.addAtEnd(12)
	li.print()
}
