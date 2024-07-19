package main

import (
	"fmt"
	"math"
)

type rectangle struct {
	l, b float64
}

func (r *rectangle) area() float64 {
	return r.l * r.b
}

func (r *rectangle) perim() float64 {
	return 2 * (r.l + r.b)
}

type circle struct {
	radius float64
}

func (c *circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c *circle) perim() float64 {
	return math.Pi * c.radius * 2
}

type geometry interface {
	area() float64
    perim() float64
}

func outArea(s geometry) {
	fmt.Println(s.area())
	fmt.Println(s.perim())
}

func main() {
	r := rectangle{1.2, 2}
	outArea(&r)
	c := circle{2}
	outArea(&c)
}
