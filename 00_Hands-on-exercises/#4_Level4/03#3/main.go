package main

import (
	"fmt"
)

func main(){
	sl := []int{42,43,44,45,46,47,48,49,50,51}

	//SLICING
	fmt.Println(sl[:5]) // 42 43 44 45 46
	fmt.Println(sl[5:])	// 47 48 49 50 51
	fmt.Println(sl[2:7]) // 44 45 46 47 48
	fmt.Println(sl[1:6]) // 43 44 45 46 47
}