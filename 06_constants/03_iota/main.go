package main 

import (
	"fmt"
)

const (
	A = iota	// 0
	B = iota	// 1
	C = iota	// 2
)

const (
	D = iota	// 0
	E			// 1
	F			// 2
)

const (
	G = iota			// 0
	H = iota * 10		// 10	
	I = iota * 10		// 20
)

func main(){
	fmt.Println(A)
	fmt.Println(B)
	fmt.Println(C)
	fmt.Println(D)
	fmt.Println(E)
	fmt.Println(F)
	fmt.Println(G)
	fmt.Println(H)
	fmt.Println(I)
}