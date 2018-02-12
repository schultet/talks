package main

import (
	"fmt"
	"time"
)

func say(what string, after int, c chan int) {
	var v int
	for {
		select { // HL
		case <-time.After(200 * time.Millisecond): // HL
			fmt.Println(what, "waiting...")
		case v = <-c: // HL
			fmt.Println(what, v)
			v++
		case c <- v: // HL
			v++
			time.Sleep(time.Duration(after) * time.Millisecond)
		}
	}
}

func main() {
	c := make(chan int)
	go say("Marco", 100, c)
	go say("Polo", 500, c)
	c <- 0
	time.Sleep(5 * time.Second)
}
