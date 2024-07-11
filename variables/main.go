package main

import "fmt"

func main() {
    // three ways to initilize variables in go
    // var variable_name type = stuff
    // var variable_name = stuff // go can deduce the type
    // variable_name := stuff // shorthand declaration if you dont want to type var
    // inilitization of variables needs type `var variable_name type`
	var a = 12
	var b int = 121

	c := 12

	var x, y, z int = 12, 12, 12
	var g int
	d, e, f := 12, 12, 12
	fmt.Print(a, b, c, d, e, f, g, x, y, z)
}
