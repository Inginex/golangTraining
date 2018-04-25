package main

import (
	"fmt"
	"os"
)

func main() {
	var c = make(chan int)
	inc := incrementer(c)
	puller(inc)

	var input string
	fmt.Scanln(&input)
}

func incrementer(c chan int) <-chan int {
	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		close(c)
	}()

	return c
}

func puller(c <-chan int) {
	go func() { // this is not necessary
		for n := range c {
			fmt.Fprintf(os.Stdout, "Channel value: %d\n", n)
		}
	}()
}
