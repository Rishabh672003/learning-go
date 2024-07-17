package main

import (
	"fmt"
)

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func rotateArray(arr []int) {
	first := arr[0]
	copy(arr, arr[1:])
	arr[len(arr)-1] = first
}

func removeDups(arrstr []string) []string {
	if len(arrstr) == 0 {
		return arrstr
	}
	result := arrstr[:1]

	for i := 1; i < len(arrstr); i++ {
		if arrstr[i] != arrstr[i-1] {
			result = append(result, arrstr[i])
		}
	}
	return result
}

func revereArray(s *[5]int) [5]int {
	for i := 0; i < int(len(s))/2; i++ {
		ends := len(s) - i - 1
		s[i], s[ends] = s[ends], s[i]
	}
	return *s
}

func main() {
	a := []int{1, 2, 3, 4, 5}
	rotateArray(a)
	b := []string{"fas", "qwq", "qwq"}
	fmt.Println(removeDups(b))
	// fmt.Println(a)
}
