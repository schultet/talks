package main

import "fmt"

type A struct{}

func (a A) Foo() { fmt.Println("Foo on A") }
func (a A) Bar() { fmt.Println("Bar on A") }

type B struct {
	A // HL
}

func (b B) Bar() { fmt.Println("Bar on B") }

func main() {
	b := B{}
	b.Foo() // HL
	b.Bar()
}
