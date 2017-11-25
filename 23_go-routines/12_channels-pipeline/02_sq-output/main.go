package main

import (
	"fmt"
)

func main() {
	for n := range sq(sq(sq(gen(2, 3)))) {
		fmt.Println(n) //256 then 6561
	}
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
