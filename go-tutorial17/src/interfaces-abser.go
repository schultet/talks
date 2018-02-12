package main

import "fmt"
import "math"

type Abser interface {
	Abs() float64
}

type Vertex struct{ x, y float64 }

func (v Vertex) Abs() float64 { return math.Sqrt(v.x*v.x + v.y*v.y) }

func main() {
	var a Abser                      // HL
	a = Vertex{3, 5}                 // HL
	fmt.Printf("Abs: %f\n", a.Abs()) // HL
}
