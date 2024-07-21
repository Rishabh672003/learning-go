package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	defer func() {
		fmt.Println(time.Since(now))
	}()

	smokeSignal := make(chan bool)
	evilninja :=[]string{"tommy", "johnny"} 
	for _, nin := range evilninja {
		go attack(nin, smokeSignal)
		fmt.Println(<-smokeSignal)
	}
}

func attack(s string, ss chan bool) {
	time.Sleep(time.Second)
	fmt.Println("throwing ninja stars at ", s)
	ss <- true
}
