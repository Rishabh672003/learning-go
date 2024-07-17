package main

import (
	"fmt"
	"time"
)

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from)
	}
}

func main2() {
	f("hello")
	go f("world")

	go func(msg string) {
		fmt.Println(msg)
	}("hello")
	time.Sleep(time.Second)
}
