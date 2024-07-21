package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"sort"
)

type employee struct {
	id   int
	name string
}

type company []employee

func (c company) Len() int {
	return len(c)
}

func (c company) Less(i, j int) bool {
	return c[i].id < c[j].id
}

func (c company) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

func main() {
	var w io.Writer
	fmt.Printf("%T : %v\n", w, w)
	w = os.Stdout
	w.Write([]byte("hello")) // "hello"
	fmt.Printf("%T : %v\n", w, w)
	w = new(bytes.Buffer)
	fmt.Printf("%T : %v\n", w, w)
	w = nil
	fmt.Printf("%T : %v", w, w)

	var buf *bytes.Buffer

	buf = new(bytes.Buffer)

	fmt.Println(buf)
	com := company{{2, "rj"}, {1, "pj"}}
	sort.Sort(com)
	fmt.Println(com)

}
