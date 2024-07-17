package main

import (
	"fmt"
	"math"
)

func pointer() {
	x := 1
	p := &x
	// p, of type *int, points to x
	fmt.Println(*p) // "1"
	*p = 2
	// equivalent to x = 2
	fmt.Println(x) // "2"
}

func pointer_test() {
	var a = new(int)
	*a = 221
	fmt.Println(*a)

	var b = &a
	**b = 12
	fmt.Println(**b)
	fmt.Println(*a)
}

func make_test() {
	var a = make([]int, 2)
	var b = &a
	fmt.Println(b)
}

func main() {
	// pointer()
	// pointer_test()
	make_test()
	a := 12
	fmt.Print(a)
	var x float32 = math.Pi
	var y float64 = math.Pi
	var z complex128 = math.Pi
}
