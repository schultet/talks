package main

import (
	"fmt"
	"math"
)

func compose(f, g func(float64) float64) func(float64) float64 {
	return func(x float64) float64 {
		return f(g(x))
	}
}

func main() {
	sincos := compose(math.Sin, math.Cos) // HL
	fmt.Printf("%T: 0.5 -> %v\n", sincos, sincos(0.5))
}
