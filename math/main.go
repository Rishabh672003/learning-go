package main

const PI = 3.141592653589793
const Google = 1e100

func main() {
	// math.Abs(-12)
	// math.Min(12, 1212)
	// fmt.Println(math.NaN())
	// fmt.Println(math.Nextafter(0.1,0.2))

	if 0.1+0.2 != 0.3 {
		panic("not equal")
	}
	var a, b, c float64 = 0.1, 0.2, 0.3

	if a+b == c {
		panic("equal")
	}

	// fmt.Println(math.Sin(3 * math.Pi / 2))
	// fmt.Println(math.Pow10(100))
	// var st = "!@#!#!#rishabh"
	//
	// for _, i := range st {
	// 	if unicode.IsLetter(i) {
	// 		fmt.Printf("%c\n", i)
	// 	}
	// }
	// fmt.Println(time.Now().Weekday())
	// var twoD [2][3]int
	// var s []string
	// s = append(s, "!2")
	// fmt.Println(s, len(s))

	// string := "rishabh"
	// fmt.Print(string[::2])
}
