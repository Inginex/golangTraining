package main

import (
	"fmt"
)

func main(){
	slc := []int{2,4,5,8,9,10,15,26,25,40}

	for i, v := range slc {
		fmt.Println(i, v)
	}
	fmt.Printf("%T", slc)
}