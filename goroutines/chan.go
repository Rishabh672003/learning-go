package main

import (
	"fmt"
)

func chan() {
	chan1 := make(chan string, 2)
	chan2 := make(chan string, 2)

	go func() {
		chan1 <- "hello\n"
		chan1 <- "world\n"
	}()

	go func() {
		chan2 <- "hello\n"
		chan2 <- "world\n"
	}()

	msgs1, msg2 := <-chan1
	fmt.Print("from chan 1\n", msgs1, msg2)
	msgs2 := <-chan2
	fmt.Print("from chan 2\n", msgs2)
	// time.Sleep(time.Second)
}
