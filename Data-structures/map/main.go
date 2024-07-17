package main

import "fmt"

func main() {
	count_map := make(map[int]int)
	arr := []int{1, 2, 3, 3, 3, 3, 2, 2, 1, 1, 2}
	for _, elem := range arr {
		count_map[elem]++
	}
	fmt.Println(len(count_map))
	for key, value := range count_map {
		fmt.Println(key, " ", value)
	}

	hashMap := make(map[int]bool)
	for _, elem := range arr {
		if !hashMap[elem] {
			hashMap[elem] = true
		}
	}
	fmt.Printf("hashMap: %v\n", hashMap)
}
