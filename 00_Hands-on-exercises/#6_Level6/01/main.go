package main

import "fmt"

func main() {
	foo := func() int {
		return 5
	}()
	barInt, barS := func() (int, string) {
		return 10, "baaaaaaaaaar!"
	}()

	fmt.Printf("%v, %v, %s ", foo, barInt, barS)

}
