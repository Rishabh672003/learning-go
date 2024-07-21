package main

import (
	"fmt"
	"learning_go/fibo"
	"sync"
)

func fiboChan(n int, st *sync.WaitGroup) {
	res := fibo.FiboRecursion(n)
	fmt.Println("from 1", res)
	st.Done()
}

var dp map[int]int = map[int]int{0: 1, 1: 1}

func fiboRecurDpchan(n int, dp map[int]int, st *sync.WaitGroup) {
	res := fibo.FiboMemo(n, dp)
	fmt.Println("from 2", res)
	st.Done()
}

func fibotabchan(n int, st *sync.WaitGroup) {
	res := fibo.FiboTab(n)
	fmt.Println("from 3", res)
	st.Done()
}

func fibospacechan(n int, st *sync.WaitGroup) {
	res := fibo.FiboSpaceOptimization(n)
	fmt.Println("from 4", res)
	st.Done()
}

func main() {
	const n = 40
	var stopper sync.WaitGroup

	stopper.Add(4)

	go fiboChan(n, &stopper)
	go fiboRecurDpchan(n, dp, &stopper)
	go fibotabchan(n, &stopper)
	go fibospacechan(n, &stopper)

	stopper.Wait()
}
