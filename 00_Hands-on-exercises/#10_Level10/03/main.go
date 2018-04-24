package main

import (
	"fmt"
)

func main() {
	c := gen()
	receive(c)

	fmt.Println("about to exit")
}

func gen() <-chan int {
	c := make(chan int)
	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		close(c)
	}()
	return c
}

func receive(c <-chan int) {
	var sum int
	out := make(chan int)
	go func() {
		for n := range c {
			sum += n
		}
		out <- sum
		close(out)
	}()
	fmt.Println(<-out)
}
