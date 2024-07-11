package main

import (
	"fmt"
	"unicode/utf8"
)

func insertionSort(arr []int) []int {
	for i := 1; i < len(arr); i++ {
		for j := i - 1; j >= 0; j-- {
			if arr[j+1] < arr[j] {
				temp := arr[j+1]
				arr[j+1] = arr[j]
				arr[j] = temp
			} else {
				break
			}
		}
	}
	return arr
}

func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}

func main() {
	s := "Hello, 京"
	fmt.Println(len(s))
	// "10"
	fmt.Println(utf8.RuneCountInString(s)) // "8"
	for i := 0; i < len(s); {
		r, size := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%d\t%c\n", i, r)
		i += size
	}
	// a := []int{1, 2, 1, 3}
	// fmt.Println(insertionSort(a))
	fmt.Println(string(0x4ea3))
	fmt.Println(comma("1000000000000000000"))
}
