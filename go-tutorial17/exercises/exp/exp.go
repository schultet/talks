package main

import (
	"fmt"
)

// Exp computes the exponential function e^x
//        ∞
// e^x =  Σ (x^n)/n!
//       n=0
func Exp(x, epsilon float64) float64 {
	// TODO: IMPLEMENT
}

func main() {
	for i := float64(0); i < 10; i++ {
		fmt.Println(Exp(i, 0.0001))
	}
}
