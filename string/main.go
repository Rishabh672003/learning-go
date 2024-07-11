package main

import (
	"fmt"
)

type fraction struct {
	num, denom float64
}

func (frac fraction) Print() string {
	return fmt.Sprintf("%g/%g", frac.num, frac.denom)
}

func main() {
	// x := "12121"
	y := fraction{1.2, 2}
	z := &y
	fmt.Println(z.num)

	dog := struct {
		name string
	}{
		"rox",
	}

	fmt.Println(dog.name)
}
