package main

import (
	"fmt"
)

func main() {
	n := average(2, 5, 10, 25, 55, 67, 86)
	fmt.Println(n)
}

func average(nums ...float64) float64 {
	var total float64
	for _, v := range nums {
		total += v
	}

	return total / float64(len(nums))
}
