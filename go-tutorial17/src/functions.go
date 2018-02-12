package main

import "fmt"

func greet(name string) { // HL
	fmt.Println("Hello,", name+"!")
} // HL

func main() { // HL
	greet("Gopher")
} // HL
