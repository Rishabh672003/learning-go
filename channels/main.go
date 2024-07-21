package main

import "fmt"

func main() {
	channel := make(chan string, 2)
	channel <- "hello"
	channel <- "world"

    for _, elem := range channel{
        fmt.Printf("%r", elem)
    }
}
