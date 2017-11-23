package main

import (
	"fmt"
)

func main(){
	//Casting
	rem := 7.25
	fmt.Printf("%T\n", rem)
	fmt.Printf("%T\n", int(rem))

	//Assertion
	var val interface{} = 7
	fmt.Printf("%T\n", val)
	fmt.Printf("%T\n", val.(int))
}