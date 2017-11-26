package main

import (
	"fmt"
)

func main() {
	var count int
	c := incrementor(2)

	for n := range c {
		count++
		fmt.Println(n)
	}

	fmt.Println("Final count: ", count)
}

func incrementor(n int) chan string {
	out := make(chan string)
	done := make(chan bool)

	// PROCESS
	for i := 0; i < n; i++ {
		go func(i int) {
			// COUNTING
			for j := 0; j < 20; j++ {
				out <- fmt.Sprint("Process: ", i, " printing: ", j)
			}
			done <- true
		}(i)
	}

	go func() {
		for i := 0; i < n; i++ {
			<-done
		}
		close(out)
	}()

	return out
}

/*
CHALLENGE #1
-- Take the code from above and change it to use channels instead of wait groups and atomicity
*/
