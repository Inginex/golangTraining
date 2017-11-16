package main

import (
	"fmt"
)

func main(){
	a := 40

	fmt.Println(a)
	fmt.Println(&a)

	var b *int = &a

	fmt.Println(b)
	fmt.Println(*b)

}