package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup
var mtx sync.Mutex
var count int

func main() {
	start := time.Now()
	for i := 1; i <= 5000; i++ {
		wg.Add(2)
		go foo("BARXX", i)
		go foo("BARYY", i)
	}
	wg.Wait()
	end := time.Since(start).Seconds()
	fmt.Printf("///////////////////// Count total: %d finished in %.2fs.\n", count, end)
}

// WITH RACE CONDITION
func foo(from string, i int) {
	defer wg.Done()
	count++
	fmt.Printf("%s Goroutine on loop %d\n and count in: %d\n", from, i, count)
}
