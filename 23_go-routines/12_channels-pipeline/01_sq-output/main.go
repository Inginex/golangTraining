package main

import (
	"fmt"
)

func main() {
	c := gen(3, 4)
	out := sq(c)

	fmt.Println(<-out)
	fmt.Println(<-out)
}

func gen(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

func sq(c <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range c {
			out <- n * n
		}
		close(out)
	}()
	return out
}
