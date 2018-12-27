package main

import "fmt"

func main() {
	foo := func(nums ...int) int {
		var tot int
		for _, num := range nums {
			tot += num
		}

		return tot
	}([]int{2, 6, 8, 12, 24, 36}...)

	bar := func(nums []int) int {
		var tot int
		for _, num := range nums {
			tot += num
		}
		return tot
	}([]int{5, 10, 8, 2, 6, 25, 13})

	fmt.Println(bar, foo)
}
