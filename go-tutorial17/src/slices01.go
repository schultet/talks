package main

import "fmt"

func main() {
	primes := [5]int{2, 3, 5, 7, 11} // array of size 5 // HL

	var s []int = primes[2:4] // slice // HL
	fmt.Printf("primes: %v\ns: %v\n", primes, s)

	s[0] = 42
	fmt.Printf("primes: %v\ns: %v\n", primes, s)
}
