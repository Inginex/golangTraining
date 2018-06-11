package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 5)
	go func() {
		for c := range ch {
			// PRETEND THAT HAS A LOT OF WORK BEING DONE HERE
			fmt.Println("Received work: ", c)
			time.Sleep(time.Duration(800) * time.Millisecond)
		}
	}()

	for i := 0; i < 20; i++ {
		select {
		case ch <- i:
			fmt.Println("Send work: ", i)
		default:
			// SERVER IS OVERLOADED
			time.Sleep(time.Duration(3) * time.Second)
			fmt.Println("Work failed: ", i)
		}
	}
	close(ch)
}
