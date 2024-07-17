package main

import "fmt"

func map_fun[T any](a T) T {
	return a
}

func main() {
	fmt.Println(map_fun(12))
}
