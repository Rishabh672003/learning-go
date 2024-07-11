package main // this refers to the name of the package this go source file belongs to

import (
	"fmt"
	"os"
	"strings"
)

func printargs1() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}

func printargs2() {
	for index, arg := range os.Args {
		fmt.Printf("%v -> %v\n", index, arg)
	}
}

func printargs3() {
	fmt.Println(strings.Join(os.Args[1:], " "))

}

func main(){
	printargs3()
}
