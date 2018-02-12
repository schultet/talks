package main

import (
	"fmt"
	"time"
)

func say(what string, after int, c chan int) {
	for {
		v := <-c
		fmt.Println(what, v)
		c <- v + 1
		time.Sleep(time.Duration(after) * time.Millisecond)
	}
}

func main() {
	c := make(chan int, 1) // HL
	go say("Ana", 100, c)
	go say("Joe", 500, c)
	c <- 1
	time.Sleep(10 * time.Second)
}
