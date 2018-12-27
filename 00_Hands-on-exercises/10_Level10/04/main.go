package main

import (
	"fmt"
)

func main() {
	q := make(chan int)
	c := gen(q)

	receive(c, q)

	fmt.Println("about to exit")
}

func receive(c <-chan int, q chan int) {
	for {
		select {
		case ch := <-c:
			fmt.Printf("Chan C: %v\n", ch)
		case qh := <-q:
			fmt.Printf("Chan QQQQQQQQ: %v\n", qh)
			return
		}
	}
}

func gen(q chan<- int) <-chan int {
	c := make(chan int)

	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		close(c)
		q <- 1
	}()

	return c

}
