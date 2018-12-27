package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup
var count int

func main() {
	for i := 0; i < 100; i++ {
		wg.Add(2)
		go incrementer("FOO", i)
		go incrementer("BAR", i)
		fmt.Println("Num of goroutines:", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("Num of goroutines:", runtime.NumGoroutine())
	fmt.Println("Final counter: ", count)
}

func incrementer(call string, i int) {
	defer wg.Done()
	x := count
	runtime.Gosched()
	x++
	fmt.Printf("%s: Counter: %d, Loop: %d\n", call, x, i)
	count = x
}
