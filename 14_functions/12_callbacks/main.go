package main

import (
	"fmt"
)

func visit(numbers []int, callback func(int)) {
	for _, n := range numbers {
		callback(n)
	}
}

func main() {
	visit([]int{1, 5, 6, 9, 15}, func(n int) {
		fmt.Println(n)
	})
}
