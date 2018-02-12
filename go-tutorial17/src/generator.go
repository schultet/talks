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

func main() {
	c1 := say("Marco", 100)
	c2 := say("Polo", 500)
	for i := 0; i < 20; i++ {
		fmt.Println(<-c1)
		fmt.Println(<-c2)
	}
}
