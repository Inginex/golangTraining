package main

import (
	"fmt"
	"math"
)

type Square struct {
	side float64
}

type Circle struct {
	radius float64
}

type Shape interface {
	area() float64
}

func (s Square) area() float64 {
	return s.side * s.side
}

func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func info(z Shape) {
	fmt.Println(z)
	fmt.Println(z.area)
}

func main(){
	sq := Square{10}
	cr := Circle{5}
	info(sq)
	info(cr)
}