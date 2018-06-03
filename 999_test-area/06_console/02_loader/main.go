package main

import (
	"fmt"
	"time"
)

func main() {
	var done = make(chan int)
	go func() {
		var loader [10]rune
		for {
			for i := 0; i < 10; i++ {
				loader[i] = '■'
				fmt.Printf("\rProgress: %c", loader)
				time.Sleep(1000 * time.Millisecond)
			}
			done <- 1
		}
	}()
	<-done
	time.Sleep(1000 * time.Millisecond)
	fmt.Printf("\rDone! 					")
}
