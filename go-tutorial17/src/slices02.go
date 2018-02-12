package main

import "fmt"

func main() {
	a := make([]float32, 5)     // a = [0 0 0 0 0] len=5 cap=5
	b := make([]float32, 0, 5)  // b = []          len=0 cap=5
	c := b[1:4]                 // c = [0 0 0]     len=3 cap=4
	fmt.Printf("%v\n%v\n%v\n", a, b, c)
}
