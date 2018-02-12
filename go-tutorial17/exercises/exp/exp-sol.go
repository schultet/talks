package main

import (
	"fmt"
	//"math"
)

// Exp computes the exponential function e^x
//        ∞
// e^x =  Σ (x^n)/n!
//       n=0
func Exp(x, epsilon float64) float64 {
	term := 1.0
	sum := term
	for i := float64(0); ; i++ {
		term = term * (x / (i + 1))
		if term < epsilon {
			return sum
		}
		sum += term
	}
}

func main() {
	for i := float64(0); i < 10; i++ {
		fmt.Println(Exp(i, 0.0001))
	}
}
