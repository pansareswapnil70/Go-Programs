package main

import (
	"fmt"
	"sync"
)

type Container struct {
	Mut      sync.Mutex
	Counters map[string]int
}

func (c *Container) inc(alpha string, integer int) {
	c.Mut.Lock()
	fmt.Println(c.Counters[alpha])
	c.Counters[alpha] += integer
	fmt.Println(c.Counters[alpha])
	c.Mut.Unlock()
}
func main() {
	c := Container{
		Counters: map[string]int{"a": 100, "b": 100},
	}

	var wg sync.WaitGroup
	DoIncrement := func(alpha string, increment int) {
		defer wg.Done()
		c.inc(alpha, increment)
	}

	wg.Add(3)
	go DoIncrement("a", 100)
	go DoIncrement("b", 100)
	go DoIncrement("a", 100)
	wg.Wait()
	fmt.Println(c.Counters)
}
