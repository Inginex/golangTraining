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
	persons := map[string]person{p1.lastname: p1, p2.lastname: p2}
	fmt.Println(persons[p1.lastname])
	fmt.Println(persons[p2.lastname])
}