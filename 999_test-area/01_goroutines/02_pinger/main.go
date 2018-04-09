package main

import (
	"fmt"
	"math/rand"
	"time"
)

func ping(c chan string) {
	for {
		c <- "ping"
		rand := rand.Intn(2000)
		time.Sleep(time.Millisecond * time.Duration(rand))
	}
}
func pong(c chan string) {
	for {
		c <- "pong"
		rand := rand.Intn(700)
		time.Sleep(time.Millisecond * time.Duration(rand))
	}
}
func printer(c chan string) {
	for {
		msg := <-c
		fmt.Printf("Time: %v \n", msg)
	}
}

func main() {
	c := make(chan string)
	go ping(c)
	go pong(c)
	go printer(c)

	var input string
	fmt.Scanln(&input)
}
