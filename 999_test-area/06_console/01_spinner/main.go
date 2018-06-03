package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {
		for {
			for _, r := range `-\|/` {
				fmt.Printf("\rLoading... %c", r)
				time.Sleep(100 * time.Millisecond)
			}
		}
	}()
	time.Sleep(2 * time.Second)
	fmt.Printf("\rDone!        ")
}
