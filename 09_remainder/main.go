package main

import "fmt"

func main(){
	// x := 13 % 3 

	for i := 1; i < 100; i++ {
		if i % 2 == 1 {
			fmt.Println("Odd")
		} else {
			fmt.Println("Even")
		}
	}
}