package main

import "fmt"

func main() {
	var v []int
	v = append(v, 1, 2, 3) // HL
	fmt.Println(v)

	w := []int{4, 5, 6}
	w = append(w, v...) // HL
	fmt.Println(w)
}
