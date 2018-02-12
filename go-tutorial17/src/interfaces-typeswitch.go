package main

import (
	"fmt"
	"strconv"
	"time"
	"math" // OMIT
)

type Abser interface { Abs() float64 } // OMIT
type Vertex struct{x,y float64} // OMIT
func (v Vertex) Abs() float64 { return math.Sqrt(v.x*v.x + v.y*v.y) } // OMIT
type Foo struct{} // OMIT
// OMIT
func (f Foo) String() string { return "Foo{}" } // OMIT
// OMIT
type Stringer interface {
	String() string
}

func ToString(any interface{}) string {
	if v, ok := any.(Stringer); ok { // HL
		return v.String() // HL
	}
	switch v := any.(type) { // HL
	case Abser: return fmt.Sprintf("Abs: %v", v.Abs()) // HL
	case int: return strconv.Itoa(v) // HL
	case float32: fmt.Sprintf("%f", v) // HL
	}
	return "???"
}

func main() {
	fmt.Println(ToString(35))
	fmt.Println(ToString(Vertex{3, 5}))
	fmt.Println(ToString(2.3 + 1.7i))
	//start := time.Now()          // OMIT
	//http.Get("http://www.google.com/") // OMIT
	//fmt.Println(ToString(time.Since(start))) // HL // OMIT
	fmt.Println(ToString(time.Minute + time.Second*10)) // HL
}
