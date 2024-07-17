package main

import "fmt"

type Queue[T any] struct {
	arr []T
}
func (q *Queue[T]) print() {
    fmt.Println(q.arr)
}

func (q *Queue[T]) isEmpty() bool {
    if len(q.arr) == 0 {
        return true
    }
    return false
}

func (q *Queue[T]) enqueue(elem T) {
    q.arr = append(q.arr, elem)
}


func (q *Queue[T]) dequeue() {
    if q.isEmpty() {
        fmt.Print("queue is empty")
        return
    }
    q.arr = q.arr[1:]
}

func main() {
    qu := Queue[int64]{}
    fmt.Println(qu.isEmpty())
    qu.enqueue(12)
    qu.enqueue(12)
    qu.enqueue(12)
    qu.print()
    qu.dequeue()
    qu.dequeue()
    qu.dequeue()
    qu.dequeue()
    qu.print()
}
