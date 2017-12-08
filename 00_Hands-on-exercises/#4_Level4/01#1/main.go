package main

import (
	"fmt"
)

func main(){
	arr := [5]int{15,25,33,45,62}

	for i, v := range arr {
		fmt.Println(i ," - ", v)
	}
	fmt.Printf("%T", arr)
}