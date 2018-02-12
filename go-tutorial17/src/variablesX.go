package main

import "fmt"

var a int
var b, c string = "alice", "bob"

func main() {
	host, port := "localhost", 3035 // HL

	fmt.Println(a, b, c, host, port)
}
