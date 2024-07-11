package main

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

func main() {
	x := "Hello!"
	// for i := 0; i < len(x); i++ {
	// 	x := x[i]
	// 	if x != '!' {
	// 		x := x + 'A' - 'a'
	// 		fmt.Printf("%c", x) // "HELLO" (one letter per iteration)
	// 	}
	// }

	for _, i := range x {
		if unicode.IsLetter(i) {
			fmt.Print(string(unicode.ToUpper(i)))
		}
	}
	fmt.Println()
	fmt.Println(utf8.RuneCountInString(x))
	// fmt.Printf(x)
}
