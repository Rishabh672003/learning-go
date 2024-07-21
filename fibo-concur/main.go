package main

import (
	"fmt"
	"learning_go/fibo"
)

func fiboChan(n int, ch chan int) {
	res := fibo.FiboRecursion(n)
	fmt.Println("from 1", res)
	ch <- res
}

var dp map[int]int = map[int]int{0: 1, 1: 1}

func fiboRecurDpchan(n int, dp map[int]int, ch chan int) {
	res := fibo.FiboMemo(n, dp)
	fmt.Println("from 2", res)
	ch <- res
}

func fibotabchan(n int, ch chan int) {
	res := fibo.FiboTab(n)
	fmt.Println("from 3", res)
	ch <- res
}

func fibospacechan(n int, ch chan int) {
	res := fibo.FiboSpaceOptimization(n)
	fmt.Println("from 4", res)
	ch <- res
}

func main() {
	const n = 40
	chan1 := make(chan int, 1)
	chan2 := make(chan int, 1)
	chan3 := make(chan int, 1)
	chan4 := make(chan int, 1)

	go fiboChan(n, chan1)
	go fiboRecurDpchan(n, dp, chan2)
	go fibotabchan(n, chan3)
	go fibospacechan(n, chan4)

    <-chan1
    <-chan2
    <-chan3
    <-chan4
}
