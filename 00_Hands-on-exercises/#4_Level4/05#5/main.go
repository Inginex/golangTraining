package main

import (
	"fmt"
)

func main(){
	var y = []int{}
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	y = append(x[:3], x[6:]...) // 42 43 44 48 49 50 51
	fmt.Println(y)
}