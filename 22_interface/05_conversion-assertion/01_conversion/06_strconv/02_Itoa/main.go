package main

import (
	"fmt"
	"strconv"
)

func main(){
	// conversion: Itoa (int to string)
	s := strconv.Itoa(12)
	fmt.Println(s)
	fmt.Printf("%T", s)
}