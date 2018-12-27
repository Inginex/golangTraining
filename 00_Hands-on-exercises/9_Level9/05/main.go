package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var count int64
var wg sync.WaitGroup

func main() {
	for i := 0; i < 100; i++ {
		wg.Add(2)
		go incremeter("FOO", i)
		go incremeter("BAR", i)
	}
	wg.Wait()
	fmt.Println("Final counter: ", count)
}

func incremeter(call string, i int) {
	defer wg.Done()
	atomic.AddInt64(&count, 1)
	fmt.Printf("%s: Counter: %d, Loop: %d\n", call, count, i)
}
