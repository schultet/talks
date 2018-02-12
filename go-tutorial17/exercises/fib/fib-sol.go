package main

import "fmt"

// fibonacci is a function that returns a function that returns an int.
func fibonacci() func() int {
	x, y, z := 0, 0, 1
	return func() int {
		x, y, z = y, z, y+z
		return x
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
	// OUTPUT:
	// 0
	// 1
	// 1
	// 2
	// 3
	// 5
	// 8
	// ...
}
