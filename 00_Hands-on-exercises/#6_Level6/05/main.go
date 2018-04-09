package main

import (
	"fmt"
	"math"
)

type square struct {
	l float64
}

func (s square) area() float64 {
	return math.Pow(s.l, 2)
}

type circle struct {
	r float64
}

func (c circle) area() float64 {
	return math.Pi * (math.Pow(c.r, 2))
}

type shape interface {
	area() float64
}

func info(s shape) {
	fmt.Println(s.area())
}

func main() {
	sq := square{20}
	cl := circle{10}
	info(cl)
	info(sq)
}
