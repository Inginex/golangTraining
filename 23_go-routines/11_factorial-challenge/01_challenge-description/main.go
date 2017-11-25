package main

import (
	"fmt"
)

func main() {
	c := fact(4)

	fmt.Println("Total: ", <-c)
}

func fact(n int) chan int {
	var out = make(chan int)
	sum := 1
	go func() {
		for i := 1; i <= n; i++ {
			sum *= i
		}
		out <- sum
		close(out)
	}()
	return out
}

/*
CHALLENGE #1:
-- Use goroutines and channels to calculate factorial

CHALLENGE #2:
-- Why might you want to use goroutines and channels to calculate factorial?
---- Formulate your own answer, then post that answer to this discussion area: https://goo.gl/flGsyX
---- Read a few of the other answers at the discussion area to see the reasons of others
*/
