package main

import (
	"fmt"
	"time"
)

func f(c chan int, quit chan bool) {
	var v int
	for {
		select {
		case <-quit:
			close(c)
			return
		case c <- v:
			v++
			time.Sleep(100 * time.Millisecond)
		}
	}
}

func main() {
	c, q := make(chan int), make(chan bool)
	go func() { time.Sleep(3 * time.Second); q <- true }()
	go f(c, q)
	for v := range c {
		fmt.Println(v)
	}
}
