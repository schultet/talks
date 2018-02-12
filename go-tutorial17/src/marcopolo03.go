package main

import (
	"fmt"
	"time"
)

func say(what string, after int) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", what, i)
			time.Sleep(time.Duration(after) * time.Millisecond)
		}
	}()
	return c
}

func fanIn(in1, in2 <-chan string) <-chan string { // HL
	c := make(chan string)
	go func() {
		for {
			select { // HL
			case v := <-in1: c <- v // HL
			case v := <-in2: c <- v // HL
			} // HL
		}
	}()
	return c
}

func main() {
	c := fanIn(say("Marco", 100), say("Polo", 500)) // HL
	for i := 0; i < 20; i++ {
		fmt.Println(<-c) // HL
	}
}
