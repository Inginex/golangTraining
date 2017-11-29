package main

import "fmt"

func main() {
	c := make(chan int)

	go func(n int) {
		for i := 1; i <= n; i++ {
			c <- i
		}
		close(c)
	}(10000)

	for n := range c {
		fmt.Print(n, "\n")
	}
}
