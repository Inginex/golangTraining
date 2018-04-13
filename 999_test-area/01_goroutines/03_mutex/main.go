package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup
var count int
var mutex sync.Mutex

func main() {
	wg.Add(2)
	go foo("foo")
	go foo("bar")
	wg.Wait()
}

func foo(caller string) {
	for i := 0; i < 10; i++ {
		time.Sleep(time.Duration(rand.Intn(20)) * time.Millisecond)
		mutex.Lock()
		count++
		fmt.Printf("[%s]: %d - counter: %d\n", caller, i, count)
		mutex.Unlock()
	}
	wg.Done()
}
