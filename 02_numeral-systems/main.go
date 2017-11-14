package main

import "fmt"

func main() {
	for i:= 60; i <= 121; i++ {
		fmt.Printf("%d \t - \t %b \t - \t %#X \t - \t %q  \n", i, i, i, i)
	}
}
