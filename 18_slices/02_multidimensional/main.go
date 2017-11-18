package main

import (
	"fmt"
)

func main(){
	list := make([][]string, 0)
	p1 := make([]string, 4)
	p1[0] = "Last Name"
	p1[1] = "First Name"
	p1[2] = "100.00"
	p1[3] = "74.00"
	list = append(list, p1)

	p2 := make([]string, 4)
	p2[0] = "Last Name"
	p2[1] = "First Name"
	p2[2] = "100.00"
	p2[3] = "75.00"
	list = append(list, p2)

	p3 := make([]string, 5)
	p3[0] = "Last Name"
	p3[1] = "First Name"
	p3[2] = "95.00"
	p3[3] = "85.00"
	p3[4] = "Drop out"
	list = append(list, p3)

	fmt.Println(list)
	fmt.Println(len(list))

}