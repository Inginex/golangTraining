package main

import (
	"fmt"
)

//Package level scope
var x int = 42

func main(){
	fmt.Println(x)
	foo()
}

func foo(){
	//Block level scope
	var y int = 12

	fmt.Println(x + y)
}