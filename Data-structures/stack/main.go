package main

import "fmt"

type Stack[T any] struct {
	arr []T
}

func (s *Stack[T]) push(elem T) {
	s.arr = append(s.arr, elem)
}

func (s *Stack[T]) String() string {
	return fmt.Sprintf("%v\n", s.arr)
}

func (s *Stack[T]) pop() T {
	var result T
	if s.isEmpty() {
		fmt.Println("Stack is empty")
		return result
	}
	result = s.arr[len(s.arr)-1]
	s.arr = s.arr[:len(s.arr)-1]
	return result
}

func (s *Stack[T]) isEmpty() bool {
	if len(s.arr) == 0 {
		return true
	}
	return false
}

func (s *Stack[T]) peek() T {
	var result T
	if s.isEmpty() {
		fmt.Println("Stack is empty")
		return result
	}
	return s.arr[len(s.arr)-1]
}

func main() {
	s := Stack[float32]{}
	s.pop()
	s.peek()
	fmt.Print(s.isEmpty())
	s.push(12)
	fmt.Println(s)
	// s.print()
}
