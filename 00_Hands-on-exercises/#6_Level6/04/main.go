package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func (u person) speak() {
	fmt.Printf("My name is: %s and I'm %d years old.", u.first, u.age)
}

func main() {
	pers := person{
		first: "Samuel",
		last:  "Palmeira",
		age:   18,
	}
	fmt.Println(pers)
	pers.speak()
}
