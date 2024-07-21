package main

import (
	"fmt"
	"os"
)

func main() {
    f, err := os.ReadFile("/home/rishabh/projects/PYTHON-ML/numpy/test.py")
    if err != nil {
      panic(err)
    }
    fmt.Printf("f: %s\n", f)
}
