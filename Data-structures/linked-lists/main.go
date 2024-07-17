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

func (l *Lili[T]) isEmpty() bool {
	if l.head == nil {
		return true
	}
	return false
}

func (l *Lili[T]) addAtEnd(elem T) {
	no := &Node[T]{elem, nil}
	if l.isEmpty() {
		l.head = no
	} else {
		node := l.head

		for node.next != nil {
			node = node.next
		}
		node.next = no
	}
}

func (l *Lili[T]) addAtStart(elem T) {
	no := &Node[T]{elem, nil}
	if l.isEmpty() {
		l.head = no
	} else {
		node := l.head
		no.next = node
		l.head = no
	}
}

func (l *Lili[T]) removeAtIndex(n int) {
	if l.isEmpty() {
		return
	}
	node := l.head
	count := 0
	for node.next != nil && count != n {
		node = node.next
		count ++
	}

}

func main() {
	li := Lili[string]{}
	li.addAtEnd("121")
	li.addAtEnd("122")
	li.addAtEnd("123")
	li.addAtStart("69")
	li.addAtStart("70")
	li.print()
}
