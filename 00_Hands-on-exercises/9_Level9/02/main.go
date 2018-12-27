package main

import (
	"fmt"
)

type person struct {
}

func (p *person) speak() string {
	return "My name is Bond, James Bond!"
}

type human interface {
	speak() string
}

func saySomething(h human) {
	fmt.Println(h.speak())
}

func main() {
	p := person{}
	saySomething(&p)
}
