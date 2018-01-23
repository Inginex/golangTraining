package main

import (
	"fmt"
)

type person struct {
	firstname string
	lastname string
	flavors []string
}

func main(){
	p1 := person{
		firstname: "Samuel",
		lastname: "Palmeira",
		flavors: []string{
			"Chocolage",
			"Strawberry",
			"Lemon",
		},
	}

	p2 := person{
		firstname: "Some",
		lastname: "body",
		flavors: []string{
			"Unknown",
		},
	}

	fmt.Println(p1, p2)
}