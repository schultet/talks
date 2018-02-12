package main

import "fmt"

func main() {
	// make returns a map of the given type, initialized and ready to use
	m := make(map[int]int)
	m[0] = 1
	m[1] = 2

	v, ok := m[1]
	if ok {
		fmt.Println(v)
	}

	delete(m, 1)
	fmt.Println(m)
}
