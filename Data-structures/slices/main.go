package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"sort"
	"strings"
)

type point struct {
	num, denom float64
}

func (p* point) Print(){
	fmt.Print(p.num, p.denom)
}

func main() {
	a := &point{num: 12.1, denom: 12.12}
	a.Print()

	var s []string
	fmt.Println("uninit:", s, s == nil, len(s) == 0)

	s = make([]string, 3)
	fmt.Println("emp:", s, "len:", len(s), "cap:", cap(s))

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	fmt.Println("len:", len(s))

	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s)

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", c)

	l := s[2:5]
	fmt.Println("sl1:", l)

	l = s[:5]
	fmt.Println("sl2:", l)

	l = s[2:]
	fmt.Println("sl3:", l)

	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)

	t2 := []string{"g", "h", "i"}
	if slices.Equal(t, t2) {
		fmt.Println("t == t2")
	}

	twoD := make([][]int, 3, 10)
	// for i := 0; i < 3; i++ {
	// 	innerLen := i + 1
	// 	twoD[i] = make([]int, innerLen)
	// 	for j := 0; j < innerLen; j++ {
	// 		twoD[i][j] = i + j
	// 	}
	// }
	fmt.Println("2d: ", twoD)

	var str_arr = []int{1, 2, 3, 4, 51, 2, 2}
	fmt.Println(cap(twoD))
	sort.Slice(str_arr, func(i, j int) bool {
		return str_arr[i] < str_arr[j]
	})
	fmt.Println(str_arr)

	// var vector_int string
	in := bufio.NewReader(os.Stdin)
	text, _ := in.ReadString('\n')
	strings.Split(text, " ")
	fmt.Print(string(text[0]))
}
