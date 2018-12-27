package main

import "fmt"

type person struct {
	First string
	Last  string
	Age   int
}

func changeMe(p *person) {
	(*p).First = "James" // *person param
	p.Last = "Bond"      // shorthand declaration
	p.Age = 007
}

func main() {
	p1 := person{"Sam", "Palm", 20}
	fmt.Printf("Before changeMe():\n %v\n", p1)
	changeMe(&p1)
	fmt.Printf("After changeMe():\n %v\n", p1)
}
