package main

import (
	"fmt"
)

func main() {
	c := make(chan int, 2)

	go func(c chan int) {
		c <- 42
		c <- 52
	}(c)

	fmt.Println(<-c)
	fmt.Println(<-c)
}
