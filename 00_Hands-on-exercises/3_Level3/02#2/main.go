package main

import "fmt"

func main() {
	var i int
	c := make(chan int)

	go func(n int) {
		for i <= n {
			c <- i
			i++
		}
		close(c)
	}(10000)

	for n := range c {
		if n > 64 && n < 91 {
			for j := 0; j < 3; j++ {
				fmt.Printf("%#U\n", n)
			}
		}
	}
}
