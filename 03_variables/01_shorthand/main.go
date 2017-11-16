package main

import (
	"fmt"
)

func main(){
	a:= "1"
	b:= 1
	c:= 1.25
	d:= true

	var e int
	var f string
	var g float64
	var h bool

	fmt.Printf("%T \n", a)
	fmt.Printf("%T \n", b)
	fmt.Printf("%T \n", c)
	fmt.Printf("%T \n", d)

	fmt.Printf("%v \n", e)
	fmt.Printf("%v \n", f)
	fmt.Printf("%v \n", g)
	fmt.Printf("%v \n", h)
}