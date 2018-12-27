package main

import (
	"fmt"
)

type person struct {
	firstname string
	lastname string
	icecream []string
}

func main(){
	p1 := person{
		firstname: "Samuel",
		lastname: "Palmeira",
		icecream: []string{
			"Chocolate",
			"Vanilla", 
			"Lemon",
		},
	}
	p2 := person{
		firstname: "Ana",
		lastname: "Silva",
		icecream: []string{
			"Strawberry",
			"Chocolate",
			"Grape",
		},
	}

	fmt.Println(p1.firstname)
	fmt.Println(p1.lastname)
	for i, v := range p1.icecream {
		fmt.Println(i, "\t" +v)
	}

	fmt.Println(p2.firstname)
	fmt.Println(p2.lastname)
	for i, v := range p2.icecream {
		fmt.Println(i, "\t" +v)
	}
}