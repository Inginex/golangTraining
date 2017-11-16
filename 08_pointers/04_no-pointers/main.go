package main

import (
	"fmt"
)

func zero(x int){
	fmt.Println("%p\n", &x)	//address in func zero
	fmt.Println(&x)	//address in func zero
	x = 0
}

func main(){
	x := 5
	fmt.Println("%p\n", &x)	//address main
	fmt.Println(&x)	//address main
	zero(x)	
	fmt.Println(x)	//still x
}