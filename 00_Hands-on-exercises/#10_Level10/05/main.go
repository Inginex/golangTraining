package main

import "fmt"

func main() {
	var c = make(chan int)

	go func() {
		c <- 1
	}()

	n, ok := <-c
	fmt.Println(n, ok)

	close(c)

	n, ok = <-c
	fmt.Println(n, ok)
}
