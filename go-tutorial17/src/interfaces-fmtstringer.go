package main

import (
	"fmt"
	"strconv"
	"time"
	"math"
)

type Abser interface {	Abs() float64 } // OMIT
type Vertex struct{ x, y float64 } // OMIT
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
	case Abser: return fmt.Sprintf("Abs: %v", v.Abs())
	case int: // HL
		return strconv.Itoa(v)
	case float32: // HL
		//return strconv.Ftoa(v, 'g', -1)
	}
	return "???"
}

func main() {
	fmt.Println(35)
	fmt.Println(Vertex{3, 5})
	fmt.Println(2.3 + 1.7i)
	//start := time.Now()          // OMIT
	//http.Get("http://www.google.com/") // OMIT
	//fmt.Println(ToString(time.Since(start))) // HL // OMIT
	fmt.Println(time.Minute + time.Second*10)
}
