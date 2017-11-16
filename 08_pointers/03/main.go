package main

import (
	"fmt"
)

func main(){
	a := 40
	fmt.Println(a)	// 40
	fmt.Println(&a) // 0x20818a220

	var b *int = &a
	fmt.Println(b)	// 0x20818a220
	fmt.Println(*b) // 40

	*b = 50
	fmt.Println(a) //50
}