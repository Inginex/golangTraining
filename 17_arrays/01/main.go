package main

import (
	"fmt"
)

func main(){
	var arr [58]string
	for i:=65; i <= 122; i++ {
		arr[i-65] = string(i)
	}

	fmt.Println(arr)
}