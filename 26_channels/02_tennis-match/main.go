package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func init() {
	rand.Seed(time.Now().Unix())
}

func main() {
	var ball = make(chan int)
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		player("Wozniacki", ball)
		wg.Done()
	}()

	go func() {
		player("Kirilenko", ball)
		wg.Done()
	}()

	ball <- 1
	wg.Wait()
}

func player(name string, ball chan int) {
	for {
		hit, ok := <-ball
		if !ok {
			fmt.Printf("Player %s Won\n", name)
			return
		}
		n := rand.Intn(100)
		if n%3 == 0 {
			fmt.Printf("Player %s Missed\n", name)
			close(ball)
			return
		}
		fmt.Printf("Player %s Hit %d\n", name, hit)
		hit++
		ball <- hit
	}
}
