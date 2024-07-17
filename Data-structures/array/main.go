package main

import (
	"fmt"
)

type Currency int

const (
	USD Currency = iota
	EUR
	GBP
	RMB
)

func app(y ...int) string {
	return fmt.Sprintf("%T", y)
}

func main() {
	// a := [...]int{1, 2, 3}
	// fmt.Printf("%T\n", a)
	// b := []int{1, 9, 3}
	// fmt.Printf("%T\n", b)
	// sort.Slice(b, func(i, j int) bool {
	// 	return b[i] < b[j]
	// })
	//
	// fmt.Printf("b: %v\n", b)
	// symbol := [...]string{USD: "$", EUR: "9", GBP: "!", RMB: ""}
	// fmt.Printf("symbol: %v\n", symbol)
	// r := [...]int{99: -1}
	// fmt.Printf("r: %v\n", r)
	// fmt.Println(a==b)
	a := []int{1, 2, 3, 4}
	// b := &a
	c := append(a, 2, 3, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 11, 1, 1)
	fmt.Println(c)
	fmt.Println(a)
	// fmt.Println(app(1,2,2,2,2,2))
	s := []int{}
	fmt.Printf("%T, %t", s, s == nil)
	aCopy := make([]int, len(a))
	copy(aCopy, a[:2])
	fmt.Println(app(1,1,1,1,1))
}
