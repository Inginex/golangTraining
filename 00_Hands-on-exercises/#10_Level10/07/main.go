package main

import (
	"fmt"
	"os"
)

func main() {
	incrementer()
	var input string
	fmt.Scanln(&input)
	fmt.Println("About to exit.")
}

func incrementer() {
	var c = make(chan int)
	for j := 0; j < 10; j++ {
		go func() {
			for i := 0; i < 10; i++ {
				c <- i
			}
		}()
		puller(c)
	}
}

func puller(c <-chan int) {
	go func() { // this is not necessary
		for n := range c {
			fmt.Fprintf(os.Stdout, "Channel value: %d\n", n)
		}
	}()
}
