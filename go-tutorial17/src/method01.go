package main

import (
	"fmt"
	"math"
)

// package main, import "fmt" and "math"
type Vertex struct {
	x, y float64
}

func (v Vertex) Abs() float64 { // HL
	return math.Sqrt(v.x*v.x + v.y*v.y)
}

func main() {
	v := Vertex{3, 5}
	fmt.Println(v.Abs()) // HL
	fmt.Println(v)
}
