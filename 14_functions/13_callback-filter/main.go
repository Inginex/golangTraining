package main

import (
	"fmt"
)

func filter(number []int, callback func(int) bool) []int {
	xs := []int{}
	for _, n := range number {
		if callback(n) {
			xs = append(xs, n)
		}
	}

	return xs
}

func main() {
	xs := filter([]int{1, 2, 4, 6, 8}, func(i int) bool {
		return i > 3
	})
	fmt.Println(xs)
}
