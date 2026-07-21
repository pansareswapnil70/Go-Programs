package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan int, 1)
	c2 := make(chan int, 1)

	go func(c1 chan int) {
		time.Sleep(time.Second * 2)
		c1 <- 1
	}(c1)

	select {
	case msg := <-c1:
		fmt.Println(msg)
	case <-time.After(time.Second * 1):
		fmt.Println("Timeout after 1 sec")
	}

	go func(c2 chan int) {
		time.Sleep(time.Second * 1)
		c2 <- 2
	}(c2)

	select {
	case msg := <-c2:
		fmt.Println(msg)
	case <-time.After(time.Second * 2):
		fmt.Println("Timeout after 2 sec")
	}

}
