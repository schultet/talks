package main

import (
	"fmt"
	"math"
)

func main() {
	e := 2.71828183
	if y := math.Pow(e, math.Pi); y < 24 { // HL
		fmt.Printf("e^π = %v\n", y)
	} else { // HL
		fmt.Printf("%v\n", y)
	} // HL
	// y is out of scope
}
