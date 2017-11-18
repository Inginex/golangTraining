package main

import (
	"fmt"
)

func great(nums []int) int {
	var gnum int

	for _, n := range nums {
		if gnum < n {
			gnum = n
		}
	}

	return gnum
}

func main() {
	re := great([]int{25, 6, 8, 12, 20, 21})
	fmt.Println(re)
}
