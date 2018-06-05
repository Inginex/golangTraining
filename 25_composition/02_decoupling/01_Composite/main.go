package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Data struct {
	Line string
}

type Xenia struct {
	Connections int
}

func (x *Xenia) Pull(d *Data) {
	x.Connections++
	d.Line = fmt.Sprintf("New line")
}

type Pillar struct {
	Stored int
}

func (p *Pillar) Store(d *Data) {
	r := rand.Intn(5)
	if r%2 != 0 {
		fmt.Println("Failed to store.")
		return
	}
	p.Stored++
	fmt.Println("Stored: ", d.Line)
}

func init() {
	rand.Seed(time.Now().Unix())
}

func main() {
	x := Xenia{}
	p := Pillar{}
	d := Data{}

	for {
		if x.Connections == 5 {
			break
		}
		x.Pull(&d)
		p.Store(&d)
	}
	fmt.Printf("Connections: %d - Stored Data: %d\n", x.Connections, p.Stored)
}
