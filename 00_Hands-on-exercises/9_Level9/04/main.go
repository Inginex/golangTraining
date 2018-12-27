package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup
var count int
var mx sync.Mutex

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
	mx.Lock()
	count++
	fmt.Printf("%s: Counter: %d, Loop: %d\n", call, count, i)
	mx.Unlock()
}
