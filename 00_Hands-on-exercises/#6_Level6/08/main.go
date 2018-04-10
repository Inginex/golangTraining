package main

import "fmt"

func main() {
	bar := foo()
	bar()
}

func foo() func() {
	return func() {
		fmt.Println("this is BAAAAAAAAAAAR!")
	}
}
