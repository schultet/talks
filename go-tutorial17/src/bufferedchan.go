package main

import "fmt"

func main() {
	c := make(chan string, 2) // HL

	c <- "buffered"
	c <- "channel"

	fmt.Println(<-c)
	fmt.Println(<-c)
}
