package main

import "fmt"

type Rupee float64
type Dollar float64

func rtod(r* Rupee) Dollar {
	return Dollar(*r / 80)
}

func (r Rupee) String() string {
    return fmt.Sprintf("rupee : %g", r)
}

func dtor(d Dollar) Rupee {
	return Rupee(d * 80)
}

func main() {
    var a = Rupee(100)
    var b = rtod(&a)
    fmt.Println(a.String())
    fmt.Print(b)
}
