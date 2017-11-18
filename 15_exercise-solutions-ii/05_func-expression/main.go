package main

import (
	"fmt"
)

func main() {

	half := func(i int, call func(n int) bool) (int, bool) {
		r := i / 2
		if call(i) {
			return r, true
		}
		return r, false
	}

	fmt.Println(half(6, func(n int) bool {
		return n%2 == 0 //even
	}))
}
