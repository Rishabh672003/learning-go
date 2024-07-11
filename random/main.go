package main

import "fmt"

func fn() (int, int) {
	return 1, 1
}

// SumIntsOrFloats sums the values of map m. It supports both int64 and float64
// as types for map values.
func SumIntsOrFloats[K comparable, V int64 | float64](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}

func generic() {
	ints := map[string]int64{
		"first":  34,
		"second": 12,
	}

	// Initialize a map for the float values
	floats := map[string]float64{
		"first":  35.98,
		"second": 26.99,
	}
	fmt.Println(SumIntsOrFloats(floats))
	fmt.Println(SumIntsOrFloats(ints))
}

func arrs() []int {
	a := []int{1, 2, 3, 4, 5}
	a = append(a, 12)
	for i, v := range a {
		fmt.Print(i, ":", v, " ")
	}
	if len(a) != 0 {
	}
	fmt.Println()
	return a
}

func main() {
	// a, b := fn()
	// fmt.Println(a, b)
	arrs()
	generic()

	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)
}
