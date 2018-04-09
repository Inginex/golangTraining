package main

import (
	"fmt"
	"math/rand"
	"time"
)

func f(n int) {
	for i := 0; i < 10; i++ {
		fmt.Printf("Routine: %v\t - \t Index: %v\n", n, i)
		tm := time.Duration(rand.Intn(250))
		time.Sleep(time.Millisecond * tm)
	}
}

func main() {
	for i := 0; i < 10; i++ {
		go f(i)
	}
	var input string
	fmt.Scanln(&input)
}
