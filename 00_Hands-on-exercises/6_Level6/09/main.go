package main

import "fmt"

func main() {
	foo := func() { fmt.Println("this is FOOOOOOOOO !") }
	bar(foo)
}

func bar(foo func()) {
	foo()
}
