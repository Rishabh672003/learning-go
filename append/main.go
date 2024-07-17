package main

func appendInt(x []int, y int) []int {
	var z []int // making a new slice 
	zlen := len(x) + 1 // calculating the length of z we will be needing, ie 1 more than that the len
	if zlen <= cap(x) {
		// There is room to grow. Extend the slice.
		z = x[:zlen] // [1, 2, 3, _, _] "_" are the capacity
	} else {
		// There is insufficient space. Allocate a new array.
		// Grow by doubling, for amortized linear complexity.
		zcap := zlen
		if zcap < 2*len(x) {
			zcap = 2 * len(x)
		}
		z = make([]int, zlen, zcap)
		copy(z, x) // a built-in function; see text
	}
	z[len(x)] = y
	return z
}

func main() {

}
