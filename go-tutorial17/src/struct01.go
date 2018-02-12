package main

import "fmt"

type Vertex struct {
	x float64
	y float64
}

func main() {
	v1 := Vertex{1, 2}    // has type Vertex
	p := &Vertex{1, 2}    // has type *Vertex
	v2 := Vertex{x: 35.0} // y:0 is implicit
	v3 := Vertex{}        // x:0 and y:0

	fmt.Println(v1, p, v2, v3)
	v1.x = 42
	fmt.Println(v1.x)
}
