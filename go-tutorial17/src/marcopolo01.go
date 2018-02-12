package main

import "fmt"
import "time"

func say(what string, after int) {
	for {
		time.Sleep(time.Duration(after) * time.Millisecond)
		fmt.Println(what)
	}
}

func main() {
	go say("Marco", 100) // HL
	go say("Polo", 500)  // HL
	time.Sleep(3 * time.Second)
}
