package main

import (
	"fmt"
)

func main() {
	slice := []float64{10, 15, 23, 56, 64}
	n := average(slice...)
	fmt.Println(n)
}

func average(nums ...float64) float64 {
	var total float64
	for _, n := range nums {
		total += n
	}

	return total / float64(len(nums))
}