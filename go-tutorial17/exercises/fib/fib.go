package main

import "fmt"

// fibonacci is a function that returns a function that returns int.
func fibonacci() func() int {
	// TODO: IMPLEMENT
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
