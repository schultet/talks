package main

import (
	"fmt"
	"math/rand"
	"time"
)

func dine(name string, left, right chan bool) {
	for {
		<-left
		fmt.Println(name, "acquired left chopstick.")
		select {
		case <-right:
			fmt.Println(name, "acquired right chopstick.")
			fmt.Println(name, "dining...")
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
			left <- true
			fmt.Println(name, "releases left chopstick.")
			right <- true
			fmt.Println(name, "releases right chopstick.")
			fmt.Println(name, "DONE.")
			return
		case <-time.After(200 * time.Millisecond):
			fmt.Println(name, "thinking...")
			left <- true
			fmt.Println(name, "releases left chopstick.")
		}
	}
}

func main() {
	cs := []chan bool{make(chan bool, 1), make(chan bool, 1), make(chan bool, 1)}
	names := []string{"Hume", "Heidegger", "Kant"}

	for i := range cs {
		go dine(names[i], cs[i], cs[(i+1)%3])
		cs[i] <- true
	}
	time.Sleep(5 * time.Second)
}
