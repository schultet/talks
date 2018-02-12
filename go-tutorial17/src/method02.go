package main

import (
	"fmt"
	"math"
)

// package main, import "fmt" and "math"
type Vertex struct {
	x, y float64
}

func (v *Vertex) Scale(f float64) { // HL
	v.x *= f
	v.y *= f
}

func main() {
	v := Vertex{3, 5}
	v.Scale(3) // HL
	fmt.Println(v)
}
