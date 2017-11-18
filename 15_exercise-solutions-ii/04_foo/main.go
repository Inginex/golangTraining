package main

import (
	"fmt"
)

func foo(nums ...int) []int {
	sl := []int{}

	for _, n := range nums {
		sl = append(sl, n)
	}

	reverse := func(slice []int) []int {
		nslice := []int{}
		ln := len(slice) - 1
		i := 0
		for ln >= i {
			nslice = append(nslice, slice[ln])
			ln--
		}
		return nslice
	}
	
	return reverse(sl)
}

func main() {
	aSlice := []int{5, 6, 7, 8, 9}
	fmt.Println(foo(1, 2, 3, 4))
	fmt.Println(foo(aSlice...))
}
