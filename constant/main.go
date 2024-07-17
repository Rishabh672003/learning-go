package main

import (
	"fmt"
)

type Weekday int

const (
	Sunday Weekday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

type Degree float64

func main() {
	fmt.Println(Monday)
}
