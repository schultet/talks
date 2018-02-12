package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Philosopher struct {
	name      string
	chopstick chan bool
	neighbour *Philosopher
}

func (p *Philosopher) say(line string) {
	fmt.Printf("%v: "+line+"\n", p.name)
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
}

func (p *Philosopher) acquireChopsticks() {
	<-p.chopstick
	fmt.Printf("%v got his chopstick.\n", p.name)
	select {
	case <-p.neighbour.chopstick:
		fmt.Printf("%v got %v's chopstick.\n", p.name, p.neighbour.name)
		return
	case <-time.After(50 * time.Millisecond):
		p.chopstick <- true
		p.say("I am thinking.")
		p.acquireChopsticks()
	}
}

func (p *Philosopher) releaseChopsticks() {
	p.chopstick <- true
	p.neighbour.chopstick <- true
}

func (p *Philosopher) dine(announce chan *Philosopher) {
	p.say("thinking...")
	p.acquireChopsticks()
	p.say("eating...")
	p.releaseChopsticks()
	announce <- p
}

func main() {
	philosophers := []*Philosopher{
		{name: "Kant"}, {name: "Heidegger"}, {name: "Wittgenstein"},
		{name: "Locke"}, {name: "Descartes"}, {name: "Newton"},
		{name: "Hume"}, {name: "Leibniz"},
	}
	for i, p := range philosophers {
		p.neighbour = philosophers[(i+1)%len(philosophers)]
		p.chopstick = make(chan bool, 1)
		p.chopstick <- true
	}
	announce := make(chan *Philosopher)
	for _, p := range philosophers {
		go p.dine(announce)
	}
	for i := 0; i < len(philosophers); i++ {
		p := <-announce
		fmt.Printf("%v is done dining.\n", p.name)
	}
}
