package main

import (
	"fmt"
	"math/rand"
	"time"
)

// dine implements a dining philosopher. Each philosopher does some thinking,
// then tries to acquire both chopsticks. If this is successful he eats and
// releases the chopsticks. He is done. Otherwise, if he has acquired one
// chopstick, he releases it again, does some more thinking, then tries to
// reaquired the chopsticks. Forever.
func dine(name string, left, right chan bool) {
	// TODO: IMPLEMENT
}

func main() {
	cs := []chan bool{make(chan bool), make(chan bool), make(chan bool)}
	names := []string{"Hume", "Heidegger", "Kant"}

	for i := range cs {
		go dine(names[i], cs[i], cs[(i+1)%3])
		go func(i int) { cs[i] <- true }(i)
	}
	time.Sleep(5 * time.Second)
}
