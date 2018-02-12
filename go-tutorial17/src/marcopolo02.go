package main

import (
	"fmt"
	"time"
)

func say(what string, after int, c chan int) {
	for {
		v := <-c // HL
		fmt.Println(what, v)
		c <- v + 1 // HL
		time.Sleep(time.Duration(after) * time.Millisecond)
	}
}

func main() {
	c := make(chan int) // HL
	go say("Marco", 100, c)
	go say("Polo", 500, c)
	c <- 0 // HL
	time.Sleep(5 * time.Second)
}
