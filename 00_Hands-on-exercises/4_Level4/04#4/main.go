package main

import (
	"fmt"
)

func main(){
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	x = append(x, 52)
	fmt.Println(x)

	xi := []int{53, 54, 55}
	x = append(x, xi...)
	
	
	fmt.Println(x)
}