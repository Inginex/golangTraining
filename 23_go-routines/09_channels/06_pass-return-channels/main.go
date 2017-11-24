package main

import (
	"fmt"
)

func main() {
	c := incrementor()
	newC := puller(c)
	fmt.Println(<-newC)
}

func incrementor() chan int {
	out := make(chan int)
	go func() {
		for i := 0; i < 20; i++ {
			out <- i
		}
		close(out)
	}()
	return out
}

func puller(c chan int) chan int {
	out := make(chan int)
	var sum int
	go func() {
		for n := range c {
			sum += n
		}
		out <- sum
		close(out)
	}()
	return out
}
